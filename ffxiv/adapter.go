package ffxiv

import (
	"fmt"
	"github.com/uppfinnarn/amai/base"
	// "github.com/PuerkitoBio/goquery"
)

type FFXIVAdapter struct {
	base.BaseAdapter
}

func NewAdapter() *FFXIVAdapter {
	return &FFXIVAdapter{
		*base.NewAdapter(map[string]string {
			"Cookie": "ldst_touchstone=1;ldst_is_support_browser=1;ldst_visit=1",
			"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		}),
	}
}

func (a *FFXIVAdapter) GetCharacter(id string) (interface{}, error) {
	doc, err := a.GetDocument(fmt.Sprintf("http://na.finalfantasyxiv.com/lodestone/character/%s/", id))
	if err != nil {
		return nil, err
	}
	
	char := &FFXIVCharacter{}
	char.ID = id
	char.Name = doc.Find(".txt_charaname").Text()
	char.Server = "Server??"
	return char, nil
}
