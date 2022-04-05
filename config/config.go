package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	botConfig *BotConfig
)

type BotConfig struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Error reading config file")
		return err
	}

	jsonErr := json.Unmarshal(file, &botConfig)

	if jsonErr != nil {
		fmt.Println("Error parsing config file")
		return err
	}

	return nil
}
