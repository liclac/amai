package ffxiv

// A FFXIV character.
type FFXIVCharacter struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ServerName   string `json:"server_name"`
}
