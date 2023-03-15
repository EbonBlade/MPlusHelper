package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"strings"
)

// MessageHandler is a structure that holds everything needed to the main handler
type MessageHandler struct {
	Logger *zerolog.Logger
}

// Handle is a main function
func (h *MessageHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Don't process bots messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := strings.ToLower(m.Content)

	// Process only messages that start with !mplus
	if !strings.HasPrefix(msg, "!mplus") {
		return
	}

	// Command format - !mplus14FH
	// Where:
	//   !mplus - prefix
	//   14 - key level, 0 or positive integer
	//   FH - short name of a dungeon, Freehold in this example

	// TODO:
	//  - get key level
	//  - get dungeon short name
	//  - create a dungeon event
	//  - post a message with some components (select time? roles?)

	switch m.Content {
	case "ping":
		h.Logger.Printf("Got ping request")
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			fmt.Println("error sending message", err)
		}
	case "pong":
		h.Logger.Printf("Got pong request")
		_, err := s.ChannelMessageSend(m.ChannelID, "Ping!")
		if err != nil {
			fmt.Println("error sending message", err)
		}
	}
}
