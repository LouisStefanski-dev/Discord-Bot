package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
)

import (
	"DiscordBot/cmds"
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
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Error getting chnnel,", err)
	}

	if m.Author.ID == BotId {
		return
	}
	if m.Content == BotPrefix+"ping" {
		cmds.Pong(s, channel)
	}
	if m.Content == BotPrefix+"greet" {
		cmds.Greet(s, m)
	}
	if strings.Contains(strings.ToLower(m.Content), "furries") {
		cmds.Disapprove(s, channel);
	}
}

func main() {
	err := ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	fmt.Println("Press Ctrl + C to close bot.")
	<-make(chan struct{})
	return
}
