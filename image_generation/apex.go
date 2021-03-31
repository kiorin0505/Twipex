package image_generation

import (
	"Twipex_project/config"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://public-api.tracker.gg/v2/apex/standard/profile/"

type apiClient struct {
	Key      string
	Platform string
	Id       string
}

type apexLawData struct {
	Data struct {
		PlatformInfo struct {
			AvatarURL string `json:"avatarUrl"`
		} `json:"platformInfo"`
		Segments []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			Stats struct {
				Kills struct {
					Value float32 `json:"value"`
				} `json:"kills"`
				Damage struct {
					Value float64 `json:"value"`
				} `json:"Damage"`
				Rankscore struct {
					Value        float32 `json:"value"`
					Rankmetadata struct {
						Rankname string `json:"rankName"`
					} `json:"metadata"`
				} `json:"rankScore"`
				Wins struct {
					Value float32 `json:"value"`
				} `json:"season7Wins"`
			} `json:"stats"`
		} `json:"segments"`
	} `json:"data"`
}

func getApexData(platform, id string) []apexLawData {
	api := apiClient{
		Key:      config.Config.Apikey,
		Platform: platform,
		Id:       id,
	}
	url := baseURL + api.Platform + "/" + api.Id
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("TRN-Api-Key", api.Key)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return nil
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []apexLawData
	err = json.Unmarshal([]byte("["+string(bytes)+"]"), &data)
	if err != nil {
		log.Printf("file=apex.go/72 action=ummmarshal error=%v", err)
		return nil
	}
	return data

}
