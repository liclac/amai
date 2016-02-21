package ffxiv

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
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

func parseEorzeanDate(s string) (sun int, moon int, err error) {
	eorzeanDateRegex := regexp.MustCompile(`(\d+)(?:st|nd|rd|th) sun of the (\d+)(?:st|nd|rd|th) (astral|umbral) moon`)
	matches := eorzeanDateRegex.FindStringSubmatch(strings.ToLower(s))
	
	if len(matches) != 4 {
		return 0, 0, ConfusedByMarkupError(fmt.Sprintf("Can't parse Eorzean date: %s", s))
	}
	
	sun, err = strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, err
	}
	
	moon, err = strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, err
	}
	
	moon = moon + (moon - 1)
	if matches[3] == "umbral" {
		moon += 1
	}
	
	return sun, moon, nil
}

func parseGuardianName(s string) string {
	parts := strings.SplitN(s, ",", 2)
	return parts[0]
}

func parseCharacter(id string, doc *goquery.Document) (char FFXIVCharacter, err error) {
	char = FFXIVCharacter{}
	
	char.ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		return char, err
	}
	
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
				case 0:
					// https://github.com/golang/go/issues/6842
					sun, moon, err := parseEorzeanDate(e.Text())
					if err != nil {
						return false
					}
					char.BirthDay = sun
					char.BirthMonth = moon
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
