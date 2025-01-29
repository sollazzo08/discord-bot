package main

import (
	"encoding/json"
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
	"time"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func formatWeatherResponse(body []byte) string {
	var weatherData WeatherResponse
	err := json.Unmarshal(body, &weatherData)
	if err != nil {
		return "Error parsing weather data."
	}

	// Convert timestamps to human-readable time
	sunrise := time.Unix(weatherData.Sys.Sunrise, 0).Format("03:04 PM")
	sunset := time.Unix(weatherData.Sys.Sunset, 0).Format("03:04 PM")

	// Extract weather details
	weatherCondition := "Unknown"
	if len(weatherData.Weather) > 0 {
		weatherCondition = weatherData.Weather[0].Description
	}

	// Format the response as a readable message
	message := fmt.Sprintf(
		"Weather for %s, %s\n\n"+
			"Temperature: %.2f°F (Feels like %.2f°F)\n"+
			"Wind: %.1f mph, Gusts up to %.1f mph\n"+
			"Cloud Cover: %d%% (%s)\n"+
			"Humidity: %d%%\n"+
			"Pressure: %d hPa\n\n"+
			"Sunrise: %s\n"+
			"Sunset: %s\n\n",
		weatherData.Name, weatherData.Sys.Country,
		weatherData.Main.Temp, weatherData.Main.FeelsLike,
		weatherData.Wind.Speed, weatherData.Wind.Gust,
		weatherData.Clouds.All, weatherCondition,
		weatherData.Main.Humidity,
		weatherData.Main.Pressure,
		sunrise, sunset,
	)

	return message
}

// Function to send the formatted weather response to Discord
func sendWeatherResponse(s *discordgo.Session, m *discordgo.MessageCreate, body []byte) {
	formattedMessage := formatWeatherResponse(body)
	s.ChannelMessageSend(m.ChannelID, formattedMessage)
}

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
			weatherData, err := io.ReadAll(resp.Body)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error reading response: "+err.Error())
				return
			}

			// Send the formatted response to Discord
			sendWeatherResponse(s, m, weatherData)
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
