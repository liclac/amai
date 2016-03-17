package ffxiv

import (
	"fmt"
	"github.com/uppfinnarn/amai/base"
	// "github.com/PuerkitoBio/goquery"
)

// Adapter for FFXIV.
type FFXIVAdapter struct {
	base.BaseAdapter
}

// Creates a new FFXIV adapter.
func NewAdapter() *FFXIVAdapter {
	return &FFXIVAdapter{
		*base.NewAdapter(map[string]string{
			"Cookie": "ldst_touchstone=1;ldst_is_support_browser=1;ldst_visit=1",
		}),
	}
}

// Gets information about a character.
func (a *FFXIVAdapter) GetCharacter(id string, results chan interface{}, errors chan error) {
	doc, err := a.GetDocument(fmt.Sprintf("http://na.finalfantasyxiv.com/lodestone/character/%s/", id))
	if err != nil {
		errors <- err
		return
	}

	char, err := parseCharacter(id, doc)
	if err != nil {
		errors <- err
		return
	}

	results <- char
}

// Gets information about a free company.
func (a *FFXIVAdapter) GetGuild(id string, results chan interface{}, errors chan error) {
	doc, err := a.GetDocument(fmt.Sprintf("http://na.finalfantasyxiv.com/lodestone/freecompany/%s/", id))
	if err != nil {
		errors <- err
		return
	}

	fc, err := parseFreeCompany(id, doc)
	if err != nil {
		errors <- err
		return
	}

	results <- fc
}
