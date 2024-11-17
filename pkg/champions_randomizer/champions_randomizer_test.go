package champions_randomizer

import (
	"testing"
)

type mockRiotClient struct{}

func (m *mockRiotClient) GetChampionData() (map[string]interface{}, error) {
	return map[string]interface{}{
		"Aatrox": map[string]interface{}{},
		"Ahri":   map[string]interface{}{},
		"Zed":    map[string]interface{}{},
		"Kennen": map[string]interface{}{},
	}, nil
}

func TestRandomizeChampions(t *testing.T) {
	client := &mockRiotClient{}
	players := 1

	randomizer := New(client)

	blue, controlMap := randomizer.RandomizeChampions(players, nil)

	red, _ := randomizer.RandomizeChampions(players, controlMap)

	if len(blue) != players*2 || len(red) != players*2 {
		t.Errorf("Expected %d champions, got %d", players*2, len(blue))
	}
}
