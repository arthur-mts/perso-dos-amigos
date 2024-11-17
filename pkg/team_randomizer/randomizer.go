package team_randomizer

import (
	"log"
	"math/rand"
	"perso-dos-amigos/bot/pkg/riot_client"
)

type Side string

const (
	Red  Side = "RED"
	Blue      = "BLUE"
)

type ChampionInfo struct {
	Photo string `json:"photo"`
	Name  string `json:"name"`
}

var (
	logger = log.Default()
)

type Team struct {
	Color     Side           `json:"color"`
	Champions []ChampionInfo `json:"champions"`
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

func GenerateTeams(players int) []Team {
	champions, err := riot_client.GetChampionData()

	if err != nil {
		logger.Panic("Failed to get champions")
	}

	championsQuantity := len(champions)
	championsNames := getKeys(champions)

	alreadySortedChampions := make(map[int]bool)

	blueTeam := &Team{}
	blueTeam.Color = Blue

	blueChampions := make([]ChampionInfo, 0, players*2)

	for len(blueChampions) < players*2 {
		championIdx := rand.Intn(championsQuantity + 1)

		if !alreadySortedChampions[championIdx] {
			alreadySortedChampions[championIdx] = true

			championName := championsNames[championIdx]
			blueChampions = append(blueChampions, ChampionInfo{
				Name:  championName,
				Photo: getChampionPhotoUrl(championName),
			})
		}
	}

	blueTeam.Champions = blueChampions

	redTeam := &Team{}
	redTeam.Color = Red

	redChampions := make([]ChampionInfo, 0, players*2)

	for len(redChampions) < players*2 {
		championIdx := rand.Intn(championsQuantity + 1)

		if !alreadySortedChampions[championIdx] {
			alreadySortedChampions[championIdx] = true

			championName := championsNames[championIdx]
			redChampions = append(redChampions, ChampionInfo{
				Name:  championName,
				Photo: getChampionPhotoUrl(championName),
			})
		}
	}
	redTeam.Champions = redChampions

	teams := []Team{*redTeam, *blueTeam}

	return teams
}

func init() {
	logger.SetPrefix("riot_client")
}
