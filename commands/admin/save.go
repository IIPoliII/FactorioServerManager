package admin

import (
	"io"

	"github.com/bwmarrin/discordgo"
)

// SaveServer executes the save command on the server.
func SaveServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	if(!*R) {
		s.ChannelMessageSend(m.ChannelID, "Command not performed. Factorio Server is not running!")
		return
	}
	io.WriteString(*P, "/save\n")
	s.ChannelMessageSend(m.ChannelID, "Factorio Server saved successfully!")
	return
}
