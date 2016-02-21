package ffxiv

// A FFXIV character.
type FFXIVCharacter struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	ServerName   string `json:"server_name"`
	Race         string `json:"race"`
	Clan         string `json:"clan"`
	Gender       string `json:"gender"`
	BirthDay     int    `json:"birth_day"`
	BirthMonth   int    `json:"birth_month"`
	Guardian     string `json:"guardian"`
	GrandCompany struct {
		Name     string `json:"name"`
		Rank     int    `json:"rank"`
	}                   `json:"grand_company"`
	FreeCompany  struct {
		ID       uint64 `json:"id"`
		Name     string `json:"name"`
	}                   `json:"free_company"`
}
