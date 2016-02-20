package ffxiv

import (
	"fmt"
	"github.com/uppfinnarn/amai/adapters"
	// "github.com/PuerkitoBio/goquery"
)

type FFXIVCharacter struct {
	id string
	name string
	server string
}

func (c FFXIVCharacter) ID() string {
	return c.id
}

func (c FFXIVCharacter) Name() string {
	return c.name
}

func (c FFXIVCharacter) String() string {
	return fmt.Sprintf("%s (%s)", c.Name(), c.server)
}

type FFXIVAdapter struct {
	adapters.BaseAdapter
}

func NewAdapter() *FFXIVAdapter {
	return &FFXIVAdapter{
		*adapters.NewAdapter(map[string]string {
			"Cookie": "ldst_touchstone=1;ldst_is_support_browser=1;ldst_visit=1",
			"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		}),
	}
}

func (a *FFXIVAdapter) GetCharacter(id string) (adapters.Character, error) {
	doc, err := a.GetDocument(fmt.Sprintf("http://na.finalfantasyxiv.com/lodestone/character/%s/", id))
	if err != nil {
		return nil, err
	}
	
	char := &FFXIVCharacter{}
	char.name = doc.Find(".txt_charaname").Text()
	char.server = "Server??"
	return char, nil
}
