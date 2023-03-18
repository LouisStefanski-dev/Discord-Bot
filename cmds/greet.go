package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Greet(s *discordgo.Session, m *discordgo.MessageCreate){
	mesg := fmt.Sprintf("Hello, how are you %s?", m.Author.Username)
	_, err := s.ChannelMessageSend(m.ChannelID, mesg)
	if err != nil{
		fmt.Println("Error sending message,", err);
	}
}