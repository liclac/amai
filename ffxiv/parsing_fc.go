package ffxiv

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
	"github.com/PuerkitoBio/goquery"
)

func parseBanner(s string) (gc, server string, err error) {
	re := regexp.MustCompile(`\s*([\w\s]+)[^\(]+\(([^\)]+)\)`)
	matches := re.FindStringSubmatch(s)
	
	if len(matches) == 0 {
		return "", "", ConfusedByMarkupError("FC Banner regex didn't match")
	}
	
	return strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2]), nil
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
		
		switch key {
		case "Free Company Name«Company Tag»":
			txt := e.Find("td").Text()
			nameTagMatches := nameTagRE.FindStringSubmatch(txt)
			if len(nameTagMatches) == 0 {
				err = ConfusedByMarkupError("Can't parse FC name/tag")
				return false
			}
			fc.Name = strings.TrimSpace(nameTagMatches[1])
			fc.Tag = strings.TrimSpace(nameTagMatches[2])
		case "Formed":
		case "Active Members":
		case "Rank":
		case "Ranking":
		case "Company Slogan":
		case "Focus":
		case "Seeking":
		case "Active":
		case "Recruitment":
		case "Estate Profile":
		default:
			err = ConfusedByMarkupError(fmt.Sprintf("Unknown item: %s", key))
			return false
		}
		return true
	})
	if err != nil {
		return fc, err
	}
	
	return fc, nil
}
