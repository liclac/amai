package ffxiv

type ClassInfo struct {
	Level        int            `json:"level"`
	ExpAt        int            `json:"exp_at"`
	ExpOf        int            `json:"exp_of"`
}

// A FFXIV character.
type FFXIVCharacter struct {
	ID           uint64         `json:"id"`
	Name         string         `json:"name"`
	Title        string         `json:"title"`
	ServerName   string         `json:"server_name"`
	Race         string         `json:"race"`
	Clan         string         `json:"clan"`
	Gender       string         `json:"gender"`
	BirthDay     int            `json:"birth_day"`
	BirthMonth   int            `json:"birth_month"`
	Guardian     string         `json:"guardian"`
	GrandCompany struct {
		Name     string         `json:"name"`
		Rank     int            `json:"rank"`
	}                           `json:"grand_company"`
	FreeCompany  struct {
		ID       uint64         `json:"id"`
		Name     string         `json:"name"`
	}                           `json:"free_company"`
	Stats        map[string]int `json:"stats"`
	Classes      map[string]ClassInfo `json:"classes"`
}
