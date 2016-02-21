package ffxiv

// A FFXIV character.
type FFXIVCharacter struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	ServerName   string `json:"server_name"`
	RaceName     string `json:"race_name"`
	ClanName     string `json:"clan_name"`
	Gender       string `json:"gender"`
}
