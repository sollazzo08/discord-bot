package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	//"strings"
	"syscall"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// strSlice := strings.Split(m.Content, " ")


	// x := len(strSlice)

	// if x
	// fmt.Println(strSlice[1])

	if m.Content == "Test" {
		s.ChannelMessageSend(m.ChannelID, "Hello, it's Dave")
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")

	//Creating a new Discord session
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Error creating the discord session", err)
		return
	}

	// Register messageCreate func as a callback for MesesageCreate events
	discord.AddHandler(messageCreate)

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
