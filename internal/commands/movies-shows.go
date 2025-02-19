package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

const CHANNELID = "786329251338911775"

func formatChannelData(messages []*discordgo.Message) string {

	if len(messages) == 0 {
		log.Println("No messages retrieved.")
		return "No messages found."
	}

	log.Printf("Fetched %d messages\n", len(messages))

	//Loop through messages directly
	var formattedMessages string
	for _, msg := range messages {
		formattedMessages += fmt.Sprintf(
			"**Message ID**: %s\n"+
				"**User**: %s (%s)\n"+
				"**Timestamp**: %s\n"+
				"**Content**: %s\n\n",
			msg.ID, msg.Author.Username, msg.Author.ID, msg.Timestamp, msg.Content,
		)
	}

	log.Println("✅ Successfully formatted messages.")
	return formattedMessages
}

// We need to first read the channel called movies-shows. We will need to grab the channel ID
// Grab all users in the channel and collect them in an array. Grab User Names. We are going to build out user profiles with their movies

// Fetching the data logic
// There is a rate limit of 100 api calls for discord
// Using the ChannelMessages function from Discord I can pull up to 100 message per call.
// When I make my first parse on the channel. I can start from the first message by getting the message ID.
// We can use this as our afterID parameter, so all 100 messages after the first will be parsed.
// We then need to track the id of the last mesg we parsed so we can start the parse over again at that mesg.
// Repeat until we reach the end, end can be determined by a response of 0 from the ChannelMessages call.

func FetchChannelData(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "!parseMovies" {
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Parsing movies...")

	var lastMessageID string = "" // Start with an empty ID (Discord fetches the most recent messages)
	limit := 100
	var allChannelMessages []*discordgo.Message

	for {
		// Fetch messages before lastMessageID
		messages, err := s.ChannelMessages(CHANNELID, limit, lastMessageID, "", "")
		if err != nil {
			fmt.Println("Error fetching messages:", err)
			s.ChannelMessageSend(m.ChannelID, "Error fetching messages.")
			return
		}

		// If no messages are returned, we've reached the beginning
		if len(messages) == 0 {
			break
		}

		// Append new messages to the collection
		allChannelMessages = append(allChannelMessages, messages...)

		// Update lastMessageID to the FIRST message fetched in the batch (moving backwards)
		lastMessageID = messages[len(messages)-1].ID
	}

	// Notify user about successful parsing
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Fetched %d messages", len(allChannelMessages)))

	formatChannelData(allChannelMessages)
}
