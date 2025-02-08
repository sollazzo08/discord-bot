package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)
// When a user reacts to a the welcome message
// We assign them a role based on the emoji they reacted on
// we need to track whenever a user reacts to the message
// we to track the mesg and user and the emojis


func ReactToRoleSelection(s *discordgo.Session, m *discordgo.MessageReactionAdd,) {
// TODO Look into MessageReactionAdd, returns data for when a Message has been reacted to



	fmt.Println("hello world")
}
