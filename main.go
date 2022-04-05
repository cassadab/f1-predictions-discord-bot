package main

import (
	"fmt"

	"github.com/cassadab/f1predbot/bot"
	"github.com/cassadab/f1predbot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start(config.Token)

	<-make(chan struct{})
	return
}
