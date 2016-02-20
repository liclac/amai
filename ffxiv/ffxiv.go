package ffxiv

import (
	"fmt"
	"github.com/uppfinnarn/amai/adapters"
	"github.com/PuerkitoBio/goquery"
)

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

func (a *FFXIVAdapter) GetCharacter(id string) (*goquery.Document, error) {
	return a.GetDocument(fmt.Sprintf("http://na.finalfantasyxiv.com/lodestone/character/%s/", id))
}
