package discord

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (r *repo) SendText(channelID, data string) error {
	_, err := r.s.ChannelMessageSend(channelID, data)
	return err
}

func (r *repo) SendTextReply(channelID, data, replyToMessageID string) error {
	msg := &discordgo.MessageSend{
		Content: data,
	}

	if replyToMessageID != "" {
		msg.Reference = &discordgo.MessageReference{
			MessageID: replyToMessageID,
			ChannelID: channelID,
		}
	}

	_, err := r.s.ChannelMessageSendComplex(channelID, msg)
	return err
}

func (c *repo) SendImage(channelID, filePath string) error {
	caption := filePath[strings.LastIndex(filePath, "/")+1 : int(len(filePath)-4)]

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	msg := &discordgo.MessageSend{
		Content: caption,
		Files: []*discordgo.File{
			{
				Name:        fmt.Sprintf("%s.png", caption),
				ContentType: "image/png",
				Reader:      file,
			},
		},
	}

	_, err = c.s.ChannelMessageSendComplex(channelID, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
