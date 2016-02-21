package ffxiv

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func normalizeServerName(name string) string {
	return strings.TrimSuffix(strings.TrimPrefix(name, "("), ")")
}

func parseCharacter(id string, doc *goquery.Document) (char FFXIVCharacter, err error) {
	char = FFXIVCharacter{}
	char.ID = id
	char.Name = doc.Find(".txt_charaname").Text()
	char.ServerName = normalizeServerName(doc.Find(".txt_worldname").Text())
	return char, nil
}
