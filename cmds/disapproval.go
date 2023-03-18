package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Disapprove(s *discordgo.Session, m *discordgo.Channel){
	_, err := s.ChannelMessageSend(m.ID, "Please don't mention those in here.")
	if err != nil{
		fmt.Println("Error sending message,", err);
	}
}