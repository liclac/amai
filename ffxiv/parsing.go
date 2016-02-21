package ffxiv

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func normalizeServerName(name string) string {
	name = strings.TrimSpace(name)
	return strings.TrimSuffix(strings.TrimPrefix(name, "("), ")")
}

func parseCharacter(id string, doc *goquery.Document) (char FFXIVCharacter, err error) {
	char = FFXIVCharacter{}
	char.ID = id
	char.Name = doc.Find(".player_name_txt h2 a").Text()
	char.ServerName = normalizeServerName(doc.Find(".player_name_txt h2 span").Text())
	return char, nil
}
