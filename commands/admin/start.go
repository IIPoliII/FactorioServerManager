package admin

import (
	"github.com/bwmarrin/discordgo"
)

// StartServer sets Stopped variable to false so auto-restart function will start server again.
func StartServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	if(*R) {
		s.ChannelMessageSend(m.ChannelID, "Factorio Server is already running!")
		return
	}
	if(!*R && Stopped) {
		Stopped = false
		return
	}

	s.ChannelMessageSend(m.ChannelID, "FactoCord's running/stopped state for Factorio Server is invalid!")
}
