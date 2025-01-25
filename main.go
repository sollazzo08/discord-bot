package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// I first need to split the entire string including the command by spaces
// Then I check the length of the slice
// If len is 2 then i take the city as index 1
// If len is 3 then I take both index 1 and 2 as the city concated by a space index 1 + " " index 2
// if len is 1 then its an error
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	strSlice := strings.Split(m.Content, " ")

	if strSlice[0] == "!weather" {

		x := len(strSlice)
		if x == 1 && strSlice[0] == "!weather" {
			s.ChannelMessageSend(m.ChannelID, "Please enter a city following the !weather command -> i.e. !weather Yonkers")
		}

		if x == 2 {
			fmt.Println(strSlice[0], strSlice[1])
		}

		if x == 3 {
			fmt.Println(strSlice[0], strSlice[1], strSlice[2])
		}

		if x == 4 {
			fmt.Println(strSlice[0], strSlice[1], strSlice[2], strSlice[4])
		}
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
