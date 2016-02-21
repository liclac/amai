package ffxiv

import (
	"testing"
)

func TestNormalizeServerID(t *testing.T) {
	if normalizeServerID("(Ultros)") != "Ultros" {
		t.Fail();
	}
	
	if normalizeServerID("Ultros") != "Ultros" {
		t.Fail();
	}
}
