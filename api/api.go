package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cassadab/f1predbot/config"
)

type StandingsResponse struct {
	Discord string `json:"discord"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Score   string `json:"score"`
}

var discordErrMsg = "Unable to retrieve standings :("

func GetStandings() string {

	client := &http.Client{}
	request, _ := http.NewRequest("GET", config.ApiBaseUrl+"predictions/standings", nil)

	request.Header.Set("x-api-key", config.ApiKey)
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("API Error: " + err.Error())
		return discordErrMsg
	}

	responseJson, err := ioutil.ReadAll(response.Body)
	var data *[]StandingsResponse

	err = json.Unmarshal(responseJson, &data)

	if err != nil {
		fmt.Println("Error parsing response: " + err.Error())
		return discordErrMsg
	}

	return FormatStandings(data)
}

func FormatStandings(standings *[]StandingsResponse) string {
	var message = ""
	for i, standing := range *standings {
		message += fmt.Sprintf("%v. %v %v\n", i+1, standing.Discord, standing.Score)
	}
	return message
}
