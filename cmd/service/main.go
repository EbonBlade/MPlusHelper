package main

import (
	"fmt"
	"github.com/EbonBlade/MPlusHelper/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Loading .env file...")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
		return
	}

	fmt.Println("Initializing bot...")
	token := os.Getenv("TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(bot.MessageHandler)
	dg.Identify.Intents = discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening websocket connection to Discord,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Bot is shutting down...")

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		fmt.Println("error closing Discord session,", err)
	}
}
