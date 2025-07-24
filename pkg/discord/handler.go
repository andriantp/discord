package discord

import (
	"encoding/json"
	"fmt"
	"strings"

	"discord/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

type Payload struct {
	ChannelID string
	MessageID string
	Content   string
}

func (r *repo) handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	logger.Trace("[handler] ChannelID:", m.ChannelID)
	logger.Trace("[handler] Author   :", m.Author.ID)
	if m.Author.ID == s.State.User.ID {
		return
	}

	logger.Trace("[handler] MessageID:", m.ID)
	logger.Trace("[handler] Content  :", m.Content)

	for alias, id := range r.channelID {
		if m.ChannelID == id {
			logger.Trace("[handler] channel match:", alias)

			if alias == "channel" {
				content := strings.ReplaceAll(m.Content, "#sahabot ", "")
				payload := &Payload{
					ChannelID: m.ChannelID,
					MessageID: m.ID,
					Content:   content,
				}

				js, err := json.MarshalIndent(payload, "", " ")
				// js, err := json.Marshal(payload)
				if err != nil {
					logger.Level("error", "handler", fmt.Sprintf("[channel] error: %s", err.Error()))
					return
				}
				logger.Trace("[handler] payload   :", string(js))
				//r.msg <- js
			}
			return
		}
	}
}
