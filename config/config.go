package config

import (
	"os"
)

var (
	Token      string
	ApiBaseUrl string
	ApiKey     string
	botConfig  *BotConfig
)

type BotConfig struct {
	Token      string `json:"token"`
	ApiBaseUrl string `json:"apiBaseUrl"`
	ApiKey     string `json:"apiKey"`
}

func ReadConfig() {
	Token = os.Getenv("TOKEN")
	ApiBaseUrl = os.Getenv("API_BASE_URL")
	ApiKey = os.Getenv("API_KEY")
}
