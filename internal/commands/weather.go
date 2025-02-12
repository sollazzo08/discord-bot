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

var (
	userRequests     = make(map[string]int)
	coolDownUserList = make(map[string]time.Time)
)

func formatWeatherResponse(body []byte) string {
	var weatherData models.WeatherResponse
	err := json.Unmarshal(body, &weatherData)
	if err != nil {
		return "Error parsing weather data."
	}

	// Get the timezone offset from OpenWeather (in seconds)
	offset := weatherData.Timezone

	// Convert sunrise & sunset to local time
	sunrise := time.Unix(weatherData.Sys.Sunrise, 0).UTC().Add(time.Second * time.Duration(offset)).Format("03:04 PM")
	sunset := time.Unix(weatherData.Sys.Sunset, 0).UTC().Add(time.Second * time.Duration(offset)).Format("03:04 PM")

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

// Rate limiting user to 5 requests in one day
func trackUserRequests(m *discordgo.MessageCreate) int {
	// When the user uses the weather command we add them to a map where the key is their discord user or id and the value is the number of requests they have made.
	// we need to create a map

	// { key: userID, value: count++}
	userRequests[m.Author.ID]++

	numberOfRequests := userRequests[m.Author.ID]

	// fmt.Println("number of requests:", numberOfRequests)

	return numberOfRequests
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Split the message content into parts
	strSlice := strings.Split(m.Content, " ")
	fmt.Println("hello")
	fmt.Println(strSlice)
	// Check if the message starts with the weather command
	if strSlice[0] == "!weatherTest" {
		fmt.Println("test")
		// If no ZIP code is provided
		if len(strSlice) == 1 {
			s.ChannelMessageSend(m.ChannelID, "Please enter a 5-digit ZIP code following the !weather command, e.g., `!weather 11111`")
			return
		}

		// If a ZIP code is provided
		if len(strSlice) == 2 {
			zip := strSlice[1]
			numberOfUserRequests := trackUserRequests(m)

			if numberOfUserRequests >= 25 {

				// lets check if they are in the cooldown list
				initialTimeStamp, exists := coolDownUserList[m.Author.ID]
				if exists {
					// get remaining cooldown time
					coolDownRemaining := time.Until(initialTimeStamp)

					if coolDownRemaining > 0 {
						message := fmt.Sprintf("You're on cooldown. Try again in %v.", coolDownRemaining.Round(time.Minute))
						s.ChannelMessageSend(m.ChannelID, message)
						return
					}
					// Cooldown expired, reset user
					delete(coolDownUserList, m.Author.ID)
					userRequests[m.Author.ID] = 0

				}

				// Start a new 24-hour cooldown for the user
				coolDownUserList[m.Author.ID] = time.Now().Add(24 * time.Hour)
				s.ChannelMessageSend(m.ChannelID, "You've reached the limit. Please wait 24 hours before using `!weather` again.")
				return
			}

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
