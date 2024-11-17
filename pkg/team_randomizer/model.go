package team_randomizer

type Side string

const (
	Red  Side = "RED"
	Blue      = "BLUE"
)

type ChampionInfo struct {
	Photo string `json:"photo"`
	Name  string `json:"name"`
}

type Team struct {
	Color     Side           `json:"color"`
	Champions []ChampionInfo `json:"champions"`
	Players   []string       `json:"players"`
}
