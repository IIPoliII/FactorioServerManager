package admin

import (
	"github.com/bwmarrin/discordgo"
)

// Restart saves and restarts the server
func Restart(s *discordgo.Session, m *discordgo.MessageCreate) {
	if *R == false {
		s.ChannelMessageSend(m.ChannelID, "Factorio Server is not running!")
		return
	}
	if *R && !Stopped {
		//Call the dedicated start/stop functions.
		StopServer(s, m)
		StartServer(s, m)
		return
	}

	s.ChannelMessageSend(m.ChannelID, "FactoCord's running/stopped state for Factorio Server is invalid!")
}
