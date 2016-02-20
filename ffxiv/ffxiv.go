package ffxiv

import (
	"github.com/uppfinnarn/amai/adapters"
)

type FFXIVAdapter struct {
	adapters.BaseAdapter
}

func NewAdapter() *FFXIVAdapter {
	return &FFXIVAdapter{*adapters.NewAdapter()}
}
