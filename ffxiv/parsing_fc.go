package ffxiv

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/uppfinnarn/amai/base"
	"regexp"
	"strconv"
	"strings"
)

func parseBanner(s string) (gc, server string, err error) {
	re := regexp.MustCompile(`\s*([\w\s]+)[^\(]+\(([^\)]+)\)`)
	matches := re.FindStringSubmatch(s)

	if len(matches) == 0 {
		return "", "", ConfusedByMarkupError("FC Banner regex didn't match")
	}

	return strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2]), nil
}

func mapTitleSwitches(lis *goquery.Selection) (res map[string]bool) {
	res = make(map[string]bool)
	lis.Each(func(_ int, li *goquery.Selection) {
		state := !li.HasClass("icon_off")
		title := li.Find("img").AttrOr("title", "")
		res[title] = state
	})
	return res
}

func parseEstateAddress(s string) (plot, ward int, district string, size int, err error) {
	re := regexp.MustCompile(`\s*Plot (\d+), (\d+) Ward, ([^\(]+)\(([^\)]+)\)`)
	matches := re.FindStringSubmatch(s)

	if len(matches) == 0 {
		return 0, 0, "", 0, ConfusedByMarkupError("Can't parse estate address")
	}

	plot, err = strconv.Atoi(matches[1])
	ward, err = strconv.Atoi(matches[2])
	district = strings.TrimSpace(matches[3])

	switch matches[4] {
	case "Small":
		size = 1
	case "Medium":
		size = 2
	case "Large":
		size = 3
	default:
		err = ConfusedByMarkupError(fmt.Sprintf("Unknown estate size: %s", matches[4]))
	}

	return plot, ward, district, size, err
}

func parseFreeCompany(id string, doc *goquery.Document) (fc FFXIVFreeCompany, err error) {
	fc = FFXIVFreeCompany{}

	fc.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fc, err
	}

	banner := doc.Find(".ic_freecompany_box").Text()
	fc.GrandCompany, fc.Server, err = parseBanner(banner)
	if err != nil {
		return fc, err
	}

	nameTagRE := regexp.MustCompile(`([^«]+)\s*«([^»]+)»`)
	doc.Find(".area_inner_body tr").EachWithBreak(func(i int, e *goquery.Selection) bool {
		key := e.Find("th").Text()
		valE := e.Find("td")

		switch key {
		case "Free Company Name«Company Tag»":
			nameTagMatches := nameTagRE.FindStringSubmatch(valE.Text())
			if len(nameTagMatches) == 0 {
				err = ConfusedByMarkupError("Can't parse FC name/tag")
				return false
			}
			fc.Name = strings.TrimSpace(nameTagMatches[1])
			fc.Tag = strings.TrimSpace(nameTagMatches[2])
		case "Formed":
		case "Active Members":
			// Skipping this in favor of parsing the full member list.
		case "Rank":
			fc.Rank, err = strconv.Atoi(strings.TrimSpace(valE.Text()))
		case "Ranking":
			// Rather uninteresting, purely ephemeral information; could parse
			// this if The Feast makes it interesting, I suppose? I honestly
			// don't even understand what's graded here.
		case "Company Slogan":
			fc.Description, err = valE.Html()
			fc.Description = base.NormalizeHtml(fc.Description)
		case "Focus":
			focus := mapTitleSwitches(e.Find("li"))
			fc.Focus = FCFocus{
				RolePlaying: focus["Role-playing"],
				Leveling:    focus["Leveling"],
				Casual:      focus["Casual"],
				Hardcore:    focus["Hardcore"],
				Dungeons:    focus["Dungeons"],
				Guildhests:  focus["Guildhests"],
				Trials:      focus["Trials"],
				Raids:       focus["Raids"],
				PvP:         focus["PvP"],
			}
		case "Seeking":
			seeking := mapTitleSwitches(e.Find("li"))
			fc.Seeking = FCSeeking{
				Tank:     seeking["Tank"],
				Healer:   seeking["Healer"],
				DPS:      seeking["DPS"],
				Crafter:  seeking["Crafter"],
				Gatherer: seeking["Gatherer"],
			}
		case "Active":
			// Not sure what the values here are... I'll have a look ingame later.
		case "Recruitment":
			fc.Recruiting = strings.TrimSpace(valE.Text()) == "Open"
		case "Estate Profile":
			mb10s := valE.Find(".mb10")
			if mb10s.Length() == 0 {
				fc.Estate = FCEstate{}
				return true
			}

			mb10s.EachWithBreak(func(j int, el *goquery.Selection) bool {
				val := el.Text()

				switch j {
				case 0:
					fc.Estate.Name = strings.TrimSpace(val)
				case 1:
					fc.Estate.Plot, fc.Estate.Ward, fc.Estate.District, fc.Estate.Size, err = parseEstateAddress(val)
				case 2:
					fc.Estate.Greeting = strings.TrimSpace(val)
				default:
					err = ConfusedByMarkupError("Too many estate info items")
				}

				return err == nil
			})
		default:
			err = ConfusedByMarkupError(fmt.Sprintf("Unknown item: %s", key))
		}
		return err == nil
	})
	if err != nil {
		return fc, err
	}

	return fc, nil
}
