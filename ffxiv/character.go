package ffxiv

type FFXIVCharacter struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Server string `json:"server"`
}
