package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Bot Parameters
var (
	Token     string
	BotPrefix string
	GuildID   string

	config *configStruct
)

// Used for structure of config.json file
type configStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
	GuildID   string `json : "GuildID"`
}

// Load values from config file
func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	GuildID = config.GuildID

	return nil
}

var BotId string
var s *discordgo.Session

func Start() {
	s, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := s.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	s.AddHandler(messageHandler)

	err = s.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	if m.Content == BotPrefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
	if m.Content == BotPrefix+"greet" {
		mesg := fmt.Sprintf("Hello, how are you %s?", m.Author.Username)
		_, _ = s.ChannelMessageSend(m.ChannelID, mesg)
	}
	if m.Content == BotPrefix+"sexyPeter" {
		mesg := fmt.Sprintf("I'm Peter Griffin, from Family Guy. That's me. Hey, %s, you're sexy", m.Author.Username)
		_, _ = s.ChannelMessageSend(m.ChannelID, mesg)
	}
	if strings.Contains(strings.ToLower(m.Content), "furries") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Please don't mention those in here.")
	}
}

func main() {
	err := ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	<-make(chan struct{})
	return
}
