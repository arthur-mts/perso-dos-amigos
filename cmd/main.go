package main

import (
	"encoding/json"
	"fmt"
	"perso-dos-amigos/bot/pkg/team_randomizer"
)

func main() {
	//champions, _ := riot_client.GetChampionData()
	//fmt.Println(champions["Taric"])

	teams := team_randomizer.GenerateTeams(5)
	data, _ := json.Marshal(teams)
	fmt.Println(string(data))
}
