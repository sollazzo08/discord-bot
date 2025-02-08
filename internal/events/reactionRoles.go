package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// When a user reacts to a the welcome message
// We assign them a role based on the emoji they reacted on
// we need to track whenever a user reacts to the message
// we to track the mesg and user and the emojis

func ReactToRoleSelection(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Define the message ID for role selection reactions
	const welcomeMessageID = "1337605328996405362"

	// Exit early if message reacted to is not the welcome message. We dont want to process the entire function for every single mesg reacted to in the server
	if m.MessageID != welcomeMessageID {
		return
	}
	// key is emoji name
	// value is role id
	roleMap := map[string]string{
		"austin_think":  "1337605328996405362", // dev role id
		"victory_crown": "1307834611644104715", // gamer role id
		"warning":       "1208164123922268190", // nsfw role id
	}
	// Assigns a role to a user based on their reaction emoji, checking if the emoji corresponds to a known role.
	if roleID, exists := roleMap[m.Emoji.Name]; exists {
		err := s.GuildMemberRoleAdd(m.GuildID, m.UserID, roleID)
		if err != nil {
			fmt.Printf("Failed to assign role: %v\n", err)
		}
	}

}
