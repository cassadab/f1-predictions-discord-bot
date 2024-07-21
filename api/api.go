package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/cassadab/f1predbot/config"
)

type StandingsResponse struct {
	// Discord string `json:"discord"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Score   int    `json:"score"`
}

const discordErrMsg = "Unable to retrieve standings :("

func GetStandings() string {

	client := &http.Client{}
	request, _ := http.NewRequest("GET", config.ApiBaseUrl+"/predictions/standings", nil)

	request.Header.Set("x-api-key", config.ApiKey)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("api error: %s", err.Error())
		return discordErrMsg
	}
	defer response.Body.Close()
	responseJson, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error reading response body: %s", err.Error())
	}

	fmt.Printf("RESPONSE: %v", responseJson)
	var data []StandingsResponse

	err = json.Unmarshal(responseJson, &data)
	if err != nil {
		fmt.Printf("error parsing response: %s", err.Error())
		return discordErrMsg
	}

	return FormatStandings(data)
}

func FormatStandings(standings []StandingsResponse) string {
	var message = "**Beeg Yoshi F1 Predictions Standings**\n"
	tableHeader := fmt.Sprintf("|%5s |%-20s | %5s", "Rank", "Name", "Score|")
	lineBreak := "\n" + strings.Repeat("-", len(tableHeader)) + "\n"
	message += "```" + "\n" + lineBreak + tableHeader + lineBreak

	for i, standing := range standings {
		message += fmt.Sprintf("|%5s |%-20s | %5d|\n", strconv.FormatInt(int64(i+1), 10), standing.Name, standing.Score)
		// message += lineBreak
	}

	message += "```"
	return message
}
