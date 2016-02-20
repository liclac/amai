package ffxiv

import (
	"fmt"
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
