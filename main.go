package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Split the message content into parts
	strSlice := strings.Split(m.Content, " ")

	// Check if the message starts with the weather command
	if strSlice[0] == "!weather" {
		// If no ZIP code is provided
		if len(strSlice) == 1 {
			s.ChannelMessageSend(m.ChannelID, "Please enter a 5-digit ZIP code following the !weather command, e.g., `!weather 11111`")
			return
		}

		// If a ZIP code is provided
		if len(strSlice) == 2 {
			zip := strSlice[1]

			// Call the weather service
			resp, err := http.Get("http://localhost:8090/api/v1/weather?zip=" + zip)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error fetching weather data: "+err.Error())
				return
			}
			defer resp.Body.Close()

			// Read the response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error reading response: "+err.Error())
				return
			}

			// Send the raw JSON back to Discord
			s.ChannelMessageSend(m.ChannelID, "Weather Data:\n```json\n"+string(body)+"\n```")
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
