package ffxiv

import (
	"github.com/PuerkitoBio/goquery"
)

func parseCharacter(id string, doc *goquery.Document) (char FFXIVCharacter, err error) {
	char = FFXIVCharacter{}
	char.ID = id
	char.Name = doc.Find(".txt_charaname").Text()
	char.Server = "Server??"
	return char, nil
}
