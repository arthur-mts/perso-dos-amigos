package main

import (
	"encoding/json"
	"fmt"
	"perso-dos-amigos/bot/pkg/champions_randomizer"
	"perso-dos-amigos/bot/pkg/riot_client"
	"perso-dos-amigos/bot/pkg/team_randomizer"
)

func main() {
	championsRandomizer := champions_randomizer.New(riot_client.New())
	championsRedTeam, controlMap := championsRandomizer.RandomizeChampions(3, nil)

	championsBlueTeam, _ := championsRandomizer.RandomizeChampions(3, controlMap)

	data, _ := json.Marshal(championsRedTeam)
	fmt.Println(string(data))

	data, _ = json.Marshal(championsBlueTeam)
	fmt.Println(string(data))

	teamRandomizer := team_randomizer.New()
	players := []string{"jonata", "wanis", "loui", "davi", "karen", "gabre"}
	blueTeam, redTeam := teamRandomizer.RandomizeTeams(players, 3)

	fmt.Println(blueTeam)
	fmt.Println(redTeam)
}
