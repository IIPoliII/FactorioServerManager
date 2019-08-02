package admin

import (
	"io"

	"github.com/bwmarrin/discordgo"
)

// UnbanPlayer executes the unban command on the server.
func UnbanPlayer(s *discordgo.Session, m *discordgo.MessageCreate, arg1 string) {
	if(!*R) {
		s.ChannelMessageSend(m.ChannelID, "Command not performed. Factorio Server is not running!")
		return
	}
	io.WriteString(*P, "/unban "+arg1+"\n")
	s.ChannelMessageSend(m.ChannelID, "Player "+arg1+" unbanned!")
	return
}
