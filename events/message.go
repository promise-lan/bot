package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

var prefix = os.Getenv("PREFIX")

func Message(s *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot || !strings.HasPrefix(event.Content, prefix) {
		return
	}
	command := string(event.Content[len(prefix):])
	if len(command) < 1 {
		return
	}
	if command == "ping" {
		start, err := event.Timestamp.Parse()
		if err != nil {
			return
		}
		temp, err := s.ChannelMessageSend(event.ChannelID, "Pinging...")
		if err != nil {
			return
		}
		end, err := temp.Timestamp.Parse()
		if err != nil {
			return
		}
		gatewayPing := s.LastHeartbeatAck.Sub(s.LastHeartbeatSent).Milliseconds()
		_, _ = s.ChannelMessageEdit(event.ChannelID, temp.ID, fmt.Sprintf("Pong! (REST: %dms, WS: %dms)", end.Sub(start).Milliseconds(), gatewayPing))
	}
}
