package admin

import (
	"io"

	"github.com/bwmarrin/discordgo"
)

// StopServer saves and stops the server.
func StopServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	if *R && !Stopped {
		io.WriteString(*P, "/save\n")
		io.WriteString(*P, "/quit\n")
		s.ChannelMessageSend(m.ChannelID, "Factorio Server saved and shutting down!")
		*R = false
		Stopped = true
		return
	}
	if !*R && Stopped {
		s.ChannelMessageSend(m.ChannelID, "Factorio Server is already stopped!")
		return;
	}

	s.ChannelMessageSend(m.ChannelID, "FactoCord's running/stopped state for Factorio Server is invalid!")
}
