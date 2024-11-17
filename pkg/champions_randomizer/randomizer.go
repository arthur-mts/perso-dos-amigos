package champions_randomizer

import (
	"log"
	"math/rand"
	"perso-dos-amigos/bot/pkg/team_randomizer"
)

var (
	logger = log.Default()
)

type RiotClient interface {
	GetChampionData() (map[string]interface{}, error)
}

type ChampionsRandomizer struct {
	riotClient RiotClient
}

func getKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getChampionPhotoUrl(championName string) string {
	return "https://ddragon.leagueoflegends.com/cdn/14.22.1/img/champion/" + championName + ".png"
}

func (c *ChampionsRandomizer) RandomizeChampions(players int, beforeSortedChampions map[int]bool) ([]team_randomizer.ChampionInfo, map[int]bool) {
	champions, err := c.riotClient.GetChampionData()

	if err != nil {
		logger.Panic("Failed to get champions")
	}

	championsQuantity := len(champions)
	championsNames := getKeys(champions)

	var alreadySortedChampions map[int]bool

	if beforeSortedChampions == nil {
		alreadySortedChampions = make(map[int]bool)
	} else {
		alreadySortedChampions = beforeSortedChampions
	}

	result := make([]team_randomizer.ChampionInfo, 0, players*2)

	for len(result) < players*2 {
		championIdx := rand.Intn(championsQuantity)

		if !alreadySortedChampions[championIdx] {
			alreadySortedChampions[championIdx] = true

			championName := championsNames[championIdx]
			result = append(result, team_randomizer.ChampionInfo{
				Name:  championName,
				Photo: getChampionPhotoUrl(championName),
			})
		}
	}

	return result, alreadySortedChampions
}

func New(riotClient RiotClient) ChampionsRandomizer {
	return ChampionsRandomizer{riotClient: riotClient}
}
func init() {
	logger.SetPrefix("riot_client")
}
