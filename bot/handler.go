package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// MessageHandler responds to "ping" and "pong" messages in a channel.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "ping":
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			fmt.Println("error sending message", err)
		}
	case "pong":
		_, err := s.ChannelMessageSend(m.ChannelID, "Ping!")
		if err != nil {
			fmt.Println("error sending message", err)
		}
	}
}
