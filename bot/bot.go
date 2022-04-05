package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cassadab/f1predbot/config"
)

var (
	Id      string
	session *discordgo.Session
)

func Start() {
	fmt.Println("Creating session")
	session, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := session.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Id = user.ID

	session.AddHandler(messageHandler)

	fmt.Println("Opening connection")
	err = session.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot running")
}

func messageHandler(sesh *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if message.Author.ID == Id {
		return
	}

	if message.Content == config.Prefix+"test" {
		_, _ = sesh.ChannelMessageSend(message.ChannelID, "This is a test!")
	}
}
