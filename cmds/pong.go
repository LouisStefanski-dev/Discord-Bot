package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Pong(s *discordgo.Session, m *discordgo.Channel) {
	_, err := s.ChannelMessageSend(m.ID, "pong")
	if err != nil {
		fmt.Println("Error sending message, ", err)
	}
}
