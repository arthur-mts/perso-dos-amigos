package team_randomizer

import "math/rand"

type teamRandomizer struct {
}

func (*teamRandomizer) RandomizeTeams(players []string, teamSize int) ([]string, []string) {
	alreadySortedPlayers := make(map[int]bool)

	blueTeam := make([]string, 0, teamSize)

	for len(blueTeam) < teamSize {
		playerIdx := rand.Intn(len(players))

		if !alreadySortedPlayers[playerIdx] {
			alreadySortedPlayers[playerIdx] = true
			blueTeam = append(blueTeam, players[playerIdx])
		}

	}

	redTeam := make([]string, 0, teamSize)
	for len(redTeam) < teamSize {
		playerIdx := rand.Intn(len(players))

		if !alreadySortedPlayers[playerIdx] {
			alreadySortedPlayers[playerIdx] = true
			redTeam = append(redTeam, players[playerIdx])
		}

	}

	return blueTeam, redTeam
}

func New() teamRandomizer {
	return teamRandomizer{}
}
