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
