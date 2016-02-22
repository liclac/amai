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

func parseGrandCompanyRank(s string) (int, error) {
	switch {
	case strings.HasSuffix(s, "Private Third Class"): return 1, nil
	case strings.HasSuffix(s, "Private Second Class"): return 2, nil
	case strings.HasSuffix(s, "Private First Class"): return 3, nil
	case strings.HasSuffix(s, "Corporal"): return 4, nil
	case strings.HasSuffix(s, "Sergeant Third Class"): return 5, nil
	case strings.HasSuffix(s, "Sergeant Second Class"): return 6, nil
	case strings.HasSuffix(s, "Sergeant First Class"): return 7, nil
	case strings.HasPrefix(s, "Chief") && strings.HasSuffix(s, "Sergeant"): return 8, nil
	case strings.HasPrefix(s, "Second") && strings.HasSuffix(s, "Lieutenant"): return 9, nil
	default: return 0, ConfusedByMarkupError(fmt.Sprintf("Unknown rank name: %s", s))
	}
}

func parseFreeCompanyIDFromURL(url string) (id uint64, err error) {
	parts := strings.Split(strings.TrimSuffix(url, "/"), "/")
	idString := parts[len(parts)-1]
	id, err = strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, err
}

func parseCharacter(id string, doc *goquery.Document) (char FFXIVCharacter, err error) {
	char = FFXIVCharacter{}
	
	char.ID, err = strconv.ParseUint(id, 10, 64)
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
					char.BirthDay, char.BirthMonth, err = parseEorzeanDate(e.Text())
					if err != nil {
						return false
					}
				case 1: char.Guardian = parseGuardianName(e.Text())
				default: err = ConfusedByMarkupError("Too many items in NamedayGuardian box")
				}
				return true
			})
		case "City-state":
			// City-state isn't particularly interesting, it affects nothing in-game. I suppose
			// you could gauge how many people chose the GC in their starting city-state, but all
			// that'd tell you is "yeah, 90% of the population did" - the game recommends picking
			// the GC in your main's city-state, which at Lv20 is typically your first class.
		case "Grand Company":
			parts := strings.Split(box.Find(".txt_name").Text(), "/")
			if len(parts) != 2 {
				err = ConfusedByMarkupError("Grand Company box isn't 'GC/Rank'")
				return false
			}
			
			char.GrandCompany.Name = parts[0]
			char.GrandCompany.Rank, err = parseGrandCompanyRank(parts[1])
			if err != nil {
				return false
			}
		case "Free Company":
			link := box.Find(".txt_name a")
			char.FreeCompany.ID, err = parseFreeCompanyIDFromURL(link.AttrOr("href", ""))
			char.FreeCompany.Name = link.Text()
			if err != nil {
				return false
			}
		default:
			err = ConfusedByMarkupError(fmt.Sprintf("Unknown infobox: %s", txt))
			return false
		}
		return true
	})
	if err != nil {
		return char, err
	}
	
	char.Stats = make(map[string]int)
	doc.Find(".param_list_attributes .right").EachWithBreak(func(i int, e *goquery.Selection) bool {
		keys := []string{ "str", "dex", "vit", "int", "mnd", "pie" }
		char.Stats[keys[i]], err = strconv.Atoi(e.Text())
		return err == nil
	})
	if err != nil {
		return char, err
	}
	doc.Find(".param_list li").EachWithBreak(func(i int, e *goquery.Selection) bool {
		keys := map[string]string{
			"Accuracy": "acc", "Critical Hit Rate": "crt", "Determination": "det",
			"Defense": "def", "Parry": "par", "Magic Defense": "mdf",
			"Attack Power": "atk", "Skill Speed": "sks", "Attack Magic Potency": "apt",
			"Healing Magic Potency": "hpt", "Spell Speed": "sps",
			"Craftsmanship": "cra", "Control": "ctl",
			"Gathering": "gat", "Perception": "pcp",
		}
		elements := e.Find("span")
		if elements.Length() != 2 {
			err = ConfusedByMarkupError("Wrong number of elements in stat item")
		}
		keyElement := elements.First()
		key, ok := keys[keyElement.Text()]
		if ok {
			valElement := elements.Last()
			char.Stats[key], err = strconv.Atoi(valElement.Text())
			return err == nil
		}
		return true
	})
	if err != nil {
		return char, err
	}
	
	char.Classes = make(map[string]ClassInfo)
	currentClassKey := ""
	currentClass := ClassInfo{}
	doc.Find(".class_list td").EachWithBreak(func(i int, e *goquery.Selection) bool {
		keys := map[string]string{
			"Gladiator": "GLA", "Marauder": "MRD", "Archer": "ARC",
			"Pugilist": "PGL", "Lancer": "LNC", "Rogue": "ROG",
			"Conjurer": "CNJ", "Arcanist": "ACN", "Dark Knight": "DRK",
			"Astrologian": "AST", "Machinist": "MCH",
			"Carpenter": "CRP", "Armorer": "ARM", "Leatherworker": "LTW",
			"Alchemist": "ALC", "Blacksmith": "BSM", "Goldsmith": "GSM",
			"Weaver": "WVR", "Culinarian": "CUL", "Miner": "MIN",
			"Fisher": "FSH", "Botanist": "BOT",
		}
		
		txt := e.Text()
		if txt == "" {
			return true
		}
		
		switch i % 3 {
		case 0:
			currentClassKey, _ = keys[txt]
		case 1:
			if txt == "-" {
				currentClass.Level = 0
			} else {
				currentClass.Level, err = strconv.Atoi(txt)
			}
		case 2:
			if txt == "- / -" {
				currentClass.ExpAt = 0
				currentClass.ExpOf = 0
			} else {
				parts := strings.Split(txt, " / ")
				if len(parts) != 2 {
					err = ConfusedByMarkupError(fmt.Sprintf("Invalid level format for %s", currentClassKey))
					return false
				}
				currentClass.ExpAt, err = strconv.Atoi(parts[0])
				currentClass.ExpOf, err = strconv.Atoi(parts[1])
			}
			
			if err == nil && len(currentClassKey) > 0 {
				char.Classes[currentClassKey] = currentClass
				currentClassKey = ""
				currentClass = ClassInfo{}
			}
		}
		return err == nil
	})
	if err != nil {
		return char, err
	}
	
	return char, nil
}
