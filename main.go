package main

import (
	"github.com/cassadab/f1predbot/bot"
	"github.com/cassadab/f1predbot/config"
)

func main() {
	config.ReadConfig()

	bot.Start()

	<-make(chan struct{})
	return
}
