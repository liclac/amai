package ffxiv

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type ConfusedByMarkupError string

func (e ConfusedByMarkupError) Error() string {
	return string(e)
}

func normalizeServerName(name string) string {
	name = strings.TrimSpace(name)
	return strings.TrimSuffix(strings.TrimPrefix(name, "("), ")")
}

func parseEorzeanDate(s string) (month int, day int) {
	return 1, 2
}

func parseGuardianName(s string) string {
	parts := strings.SplitN(s, ",", 2)
	return parts[0]
}

func parseCharacter(id string, doc *goquery.Document) (char FFXIVCharacter, err error) {
	char = FFXIVCharacter{}
	
	char.ID = id
	char.Name = doc.Find(".player_name_txt h2 a").Text()
	char.Title = doc.Find(".chara_title").Text()
	char.ServerName = normalizeServerName(doc.Find(".player_name_txt h2 span").Text())
	
	infoParts := strings.Split(doc.Find(".chara_profile_title").Text(), "/")
	char.Race = strings.TrimSpace(infoParts[0])
	char.Clan = strings.TrimSpace(infoParts[1])
	char.Gender = strings.TrimSpace(infoParts[2])
	
	doc.Find(".chara_profile_box_info").EachWithBreak(func(i int, box *goquery.Selection) bool {
		txt := box.Find(".txt").Text()
		switch txt {
		case "NamedayGuardian":
			box.Find(".txt_name").EachWithBreak(func(i int, e *goquery.Selection) bool {
				switch i {
				case 0: char.BirthMonth, char.BirthDay = parseEorzeanDate(e.Text())
				case 1: char.Guardian = parseGuardianName(e.Text())
				default: err = ConfusedByMarkupError("Too many items in NamedayGuardian box")
				}
				return true
			})
		case "City-state":
		case "Grand Company":
		case "Free Company":
		default:
			err = ConfusedByMarkupError(fmt.Sprintf("Unknown infobox: %s", txt))
			return false
		}
		return true
	})
	if err != nil {
		return char, err
	}
	
	return char, nil
}
