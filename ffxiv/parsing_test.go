package ffxiv

import (
	"testing"
)

func TestNormalizeServerName(t *testing.T) {
	if normalizeServerName("(Ultros)") != "Ultros" {
		t.Fail()
	}
	
	if normalizeServerName("Ultros") != "Ultros" {
		t.Fail()
	}
}

func TestParseEorzeanDate(t *testing.T) {
	sun, moon, err := parseEorzeanDate("1st Sun of the 1th Astral Moon")
	if sun != 1 || moon != 1 || err != nil {
		t.Fatal("1/1A: ", sun, moon, err)
	}
	
	sun, moon, err = parseEorzeanDate("2nd Sun of the 1st Umbral Moon")
	if sun != 2 || moon != 2 || err != nil {
		t.Fatal("2/1U: ", sun, moon, err)
	}
	
	sun, moon, err = parseEorzeanDate("3rd Sun of the 2nd Astral Moon")
	if sun != 3 || moon != 3 || err != nil {
		t.Fatal("3/2A: ", sun, moon, err)
	}
}

func TestParseEorzeanDateInvalid(t *testing.T) {
	_, _, err := parseEorzeanDate("Mary had a little lamb")
	if err == nil {
		t.Fail()
	}
}

func TestParseGuardianName(t *testing.T) {
	if parseGuardianName("Oschon, the Wanderer") != "Oschon" {
		t.Fail()
	}
	
	if parseGuardianName("Oschon") != "Oschon" {
		t.Fail()
	}
}

func TestParseGrandCompanyRank(t *testing.T) {
	ranks := map[string]int {
		"Storm Private Third Class": 1,
		"Storm Private Second Class": 2,
		"Storm Private First Class": 3,
		"Storm Corporal": 4,
		"Storm Sergeant Third Class": 5,
		"Storm Sergeant Second Class": 6,
		"Storm Sergeant First Class": 7,
		"Chief Storm Sergeant": 8,
		"Second Storm Lieutenant": 9,
	}
	
	for r, n := range ranks {
		i, err := parseGrandCompanyRank(r)
		if i != n {
			t.Fatal(r, n, i)
		}
		if err != nil {
			t.Fatal(r, err)
		}
	}
}

func TestParseGrandCompanyRankInvalid(t *testing.T) {
	_, err := parseGrandCompanyRank("Storm Shenanigans")
	if err == nil {
		t.Fail()
	}
}
