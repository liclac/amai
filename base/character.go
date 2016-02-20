package base

import (
	"fmt"
)

type Character interface {
	fmt.Stringer
	ID() string
	Name() string
}
