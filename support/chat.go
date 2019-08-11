package support

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hpcloud/tail"
)

// Chat pipes in-game chat to Discord.
func Chat(s *discordgo.Session, m *discordgo.MessageCreate) {
	for {
		t, err := tail.TailFile("factorio.log", tail.Config{Follow: true})
		if err != nil {
			ErrorLog(fmt.Errorf("%s: An error occurred when attempting to tail factorio.log\nDetails: %s", time.Now(), err))
		}
		for line := range t.Lines {
			if !strings.Contains(line.Text, "TransmissionControlHelper.cpp") || !strings.Contains(line.Text, "New RCON connection from IP ADDR") || !strings.Contains(line.Text, "ServerRouter.cpp") {
				if !strings.Contains(line.Text, "New RCON connection from IP ADDR") {
					s.ChannelMessageSend(Config.FactorioConsoleChatID, fmt.Sprintf("%s", line.Text))
				}
			}
			if strings.Contains(line.Text, "[CHAT]") || strings.Contains(line.Text, "[EMBED]") || strings.Contains(line.Text, "[JOIN]") || strings.Contains(line.Text, "[LEAVE]") || strings.Contains(line.Text, "[KICK]") || strings.Contains(line.Text, "[BAN]") || strings.Contains(line.Text, "[JAPC-EVENT-HANDLE]") {
				if !strings.Contains(line.Text, "<server>") || Config.PassConsoleChat {

					if strings.Contains(line.Text, "[JOIN]") ||
						strings.Contains(line.Text, "[LEAVE]") {
						TmpList := strings.Split(line.Text, " ")
						// Don't hard code the channelID! }:<
						s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", strings.Join(TmpList[3:], " ")))
					} else if strings.Contains(line.Text, "[EMBED]") {
						TmpList := strings.Split(line.Text, " ")
						message := new(discordgo.MessageSend)
						err := json.Unmarshal([]byte(fmt.Sprintf("%s", strings.Join(TmpList[3:], " "))), message)
						if err == nil {
							message.Tts = false
							s.ChannelMessageSendComplex(Config.FactorioChannelID, message)
						}
					} else if strings.Contains(line.Text, "[JAPC-EVENT-HANDLE]") {
							TmpList := strings.Split(line.Text, " ")
							s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", strings.Join(TmpList[1:], " ")))
					} else {

						TmpList := strings.Split(line.Text, " ")
						TmpList[3] = strings.Replace(TmpList[3], ":", "", -1)
						if strings.Contains(strings.Join(TmpList, " "), "@") {
							index := LocateMentionPosition(TmpList)

							for _, position := range index {
								User := SearchForUser(TmpList[position])

								if User == nil {
									continue
								}
								TmpList[position] = User.Mention()
							}

						}

						s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("<%s>: %s", TmpList[3], strings.Join(TmpList[4:], " ")))
					}
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
