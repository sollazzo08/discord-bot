package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const CHANNELID = "786329251338911775"

// We want to return a general MessageData Struct, general enough so that AI can extract the correct detail
// WE NEED A SUB STRUCT FOR GETTING THE AUTHOR DETAILS OF A EMOJI REACTION ON A MESSAGE
type MessageData struct {
	ID                   string // Message ID
	AuthorID             string // User ID of the sender
	Username             string // Username of the sender
	Content              string // The actual message text
	Timestamp            string // Message timestamp
	EmojiReactionDetails struct {
		AuthorID    string // User ID of the sender
		Username    string // Username of the sender
		EmojiDetail string //Emoji details, should be a number rating

	}
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

	// Print message content to console
	for _, msg := range allChannelMessages {
		fmt.Println(msg)
	}
}
