package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sollazzo08/discord-bot/internal/models"
)

func formatWeatherResponse(body []byte) string {
	var weatherData models.WeatherResponse
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

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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
