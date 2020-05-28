package main

import (
	"os"

	"github.com/jadech32/discordbot/pkg/modules"

	"github.com/jadech32/discordbot/pkg/bootstrap"
	"github.com/jadech32/discordbot/pkg/discord"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	bootstrap.SetLogLevel(log.InfoLevel)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get Env Variables
	token := os.Getenv("DISCORD_TOKEN")
	prefix := os.Getenv("BOT_PREFIX")

	if token == "" {
		log.Fatal("DISCORD_TOKEN cannot be empty")
	}

	log.Info("Starting Discord Bot")

	// Initialize Modules and store
	md := modules.NewModules()

	// Add Modules Here
	md.AddModule("exchange", modules.Convert{})
	md.AddModule("help", modules.Help{})

	// Other handlers (e.g. reaction listeners, etc.) can be added using md.AddHandler()

	dg, err := discord.NewDiscord(token, md)
	if err != nil {
		log.Fatal(err)
	}

	dg.SetPrefix(prefix)
	dg.InitModules()
	dg.InitHandlers()

	dg.Start()
	defer dg.Close()
	c := make(chan struct{})
	<-c
}
