package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sollazzo08/discord-bot/config"
	"github.com/sollazzo08/discord-bot/internal/ai"
	"github.com/sollazzo08/discord-bot/internal/commands"
	"github.com/sollazzo08/discord-bot/internal/events"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	//Creating a new Discord session
	discord, err := discordgo.New("Bot " + cfg.BOTTOKEN)
	if err != nil {
		fmt.Println("Error creating the discord session", err)
		return
	}

	// Register messageCreate func as a callback for MesesageCreate events

	// Pass APP_ENV using a closure
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		commands.MessageCreate(s, m, cfg.APP_ENV)
	})
	discord.AddHandler(events.ReactToRoleSelection)
	discord.AddHandler(commands.FetchChannelData)

	// Pass OPEN_AI_TOKEN using a closure
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		ai.UseChatGPT4(s, m, cfg.OPEN_AI_TOKEN)
	})

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	// Keeps the program running until it receives a termination signal (e.g., CTRL-C).
	// 1. A channel named `sc` is created to listen for OS signals (like "CTRL-C").
	// 2. The `signal.Notify` function tells Go to send specific signals (SIGINT, SIGTERM, etc.) to this channel.
	// 3. The `<-sc` statement pauses the program, waiting for a signal to arrive.
	// Once a signal is received, the program continues execution (usually to clean up and exit).

	fmt.Println("Dave Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close down the Discord session.
	discord.Close()
}
