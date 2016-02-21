package ffxiv

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func normalizeServerID(id string) string {
	return strings.TrimSuffix(strings.TrimPrefix(id, "("), ")")
}

func parseCharacter(id string, doc *goquery.Document) (char FFXIVCharacter, err error) {
	char = FFXIVCharacter{}
	char.ID = id
	char.Name = doc.Find(".txt_charaname").Text()
	char.ServerID = normalizeServerID(doc.Find(".txt_worldname").Text())
	return char, nil
}
