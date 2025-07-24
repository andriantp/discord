package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type repo struct {
	s         *discordgo.Session
	channelID map[string]string

	msg chan []byte
}

type RepositoryI interface {
	Close() error
	SendText(channelID, data string) error
	SendTextReply(channelID, data, replyToMessageID string) error
	SendImage(channelID, filePath string) error

	OnChannel() chan []byte
}

func NewRepo(token string, channelID map[string]string) (RepositoryI, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("New:%w", err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentMessageContent

	r := &repo{
		channelID: channelID,
		s:         dg,
		msg:       make(chan []byte),
	}
	dg.AddHandler(r.handler)

	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("Open:%w", err)
	}

	return r, nil
}

func (r *repo) Close() error {
	return r.s.Close()
}

func (r *repo) OnChannel() chan []byte {
	return r.msg
}
