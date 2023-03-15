package main

import (
	"github.com/EbonBlade/MPlusHelper/bot"
	logger2 "github.com/EbonBlade/MPlusHelper/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := logger2.Logger()
	logger.Print("Loading .env file...")
	err := godotenv.Load(".env")
	if err != nil {
		logger.Print("error loading .env file")
		return
	}

	logger.Print("Initializing bot...")
	token := os.Getenv("TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Print("error creating Discord session,", err)
		return
	}

	messageHandler := bot.MessageHandler{Logger: logger}
	dg.AddHandler(messageHandler.Handle)
	dg.Identify.Intents = discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		logger.Print("error opening websocket connection to Discord,", err)
		return
	}

	logger.Print("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logger.Print("Bot is shutting down...")

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		logger.Print("error closing Discord session,", err)
	}
}
