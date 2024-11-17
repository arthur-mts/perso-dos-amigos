package riot_client

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

const datagragonUrl = "https://ddragon.leagueoflegends.com"

var (
	logger = log.Default()
)

type riotClient struct {
}

func init() {
	logger.SetPrefix("riot_client")
}

func getLastLeagueVersion() (string, error) {
	resp, err := http.Get(datagragonUrl + "/api/versions.json")

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("request failed with status: " + resp.Status)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Printf("Failed to close response reader: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	versionsList := make([]string, 10)

	err = json.Unmarshal(body, &versionsList)

	if err != nil {
		logger.Printf("Failed to parse response body: %s", err)
		return "", err
	}

	if len(versionsList) > 0 {
		return versionsList[0], nil
	}

	return "", errors.New("no version returned")
}

func getChampionsFileName(version string) string {
	return "/tmp/champions_" + version + ".json"
}

func getChampionDataFromFSOrAPI() (map[string]interface{}, error) {
	lastDatadragonVersion, err := getLastLeagueVersion()

	if err != nil {
		return nil, err
	}

	championsDataFileName := getChampionsFileName(lastDatadragonVersion)
	if _, err := os.Stat(championsDataFileName); err == nil {
		file, err := os.ReadFile(championsDataFileName)

		if err != nil {
			logger.Printf("Failed to load champins file: %s", err)
			return nil, err
		}

		championsData := make(map[string]interface{})
		err = json.Unmarshal(file, &championsData)

		if err != nil {
			logger.Printf("Failed to decode chanmpions file: %s", err)
		}

		//logger.Println("Returned cached file")
		return championsData, nil
	}

	resp, err := http.Get(datagragonUrl + "/cdn/" + lastDatadragonVersion + "/data/en_US/champion.json")
	//logger.Println("Downloaded json file")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed with status: " + resp.Status)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Printf("Failed to close response reader: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	parsedBody := make(map[string]interface{})

	err = json.Unmarshal(body, &parsedBody)

	if err != nil {
		return nil, err
	}

	championsData := parsedBody["data"].(map[string]interface{})

	championsDataEncoded, err := json.Marshal(championsData)
	err = os.WriteFile(championsDataFileName, championsDataEncoded, 0644)

	if err != nil {
		return nil, err
	}

	return championsData, nil
}

func (r *riotClient) GetChampionData() (map[string]interface{}, error) {
	championsData, err := getChampionDataFromFSOrAPI()

	if err != nil {
		logger.Printf("Failed to get champion data: %s", err)
		return nil, err
	}

	return championsData, nil
}

func New() *riotClient {
	return &riotClient{}
}
