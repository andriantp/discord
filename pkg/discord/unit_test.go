package discord_test

import (
	"discord/pkg/discord"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	token     = ""
	channelID = map[string]string{
		"channel": "13",
	}
	messageID = "13"
	image     = "image/image2.png"
)

func Test_Connect(t *testing.T) {
	_, err := discord.NewRepo(token, channelID)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}
}

func Test_SendText(t *testing.T) {
	repo, err := discord.NewRepo(token, channelID)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}

	if err := repo.SendText(channelID["channel"], "hi"); err != nil {
		t.Fatalf("SendText: %v", err)
	}
}

func Test_SendReply(t *testing.T) {
	repo, err := discord.NewRepo(token, channelID)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}

	if err := repo.SendTextReply(channelID["channel"], "ada apa ?", messageID); err != nil {
		t.Fatalf("SendTextReply: %v", err)
	}
}

func Test_SendImg2(t *testing.T) {
	repo, err := discord.NewRepo(token, channelID)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := filepath.Join(base, image)

	if err := repo.SendImage(channelID["channel"], path); err != nil {
		t.Fatalf("SendImage: %v", err)
	}
}

// Table-Driven Test -> unt case ini, test logic-nya bisa diabstraksikan
func Test_Discord_Drivenx3(t *testing.T) {
	repo, err := discord.NewRepo(token, channelID)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := filepath.Join(base, image)

	tests := []struct {
		name     string
		sendType string // "text" or "image"
		channel  string
		content  string
	}{
		{"send text to channel", "text", channelID["channel"], "test"},
		{"send image to channel", "image", channelID["channel"], path},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.sendType {
			case "text":
				if err := repo.SendText(tc.channel, tc.content); err != nil {
					t.Fatalf("SendText: %v", err)
				}
			case "image":
				abs, _ := filepath.Abs(tc.content)
				if err := repo.SendImage(tc.channel, abs); err != nil {
					t.Fatalf("SendImage: %v", err)
				}
			}
		})
	}
}

// subtest
func Test_SendMethodsx(t *testing.T) {
	repo, err := discord.NewRepo(token, channelID)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}

	t.Run("SendText", func(t *testing.T) {
		err := repo.SendText(channelID["channel"], "subtest")
		if err != nil {
			t.Fatalf("SendText: %v", err)
		}
	})

	t.Run("SendImage", func(t *testing.T) {
		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)
		base := basepath[0:strings.Index(basepath, "pkg")]
		path := filepath.Join(base, image)

		err := repo.SendImage(channelID["channel"], path)
		if err != nil {
			t.Fatalf("SendImage: %v", err)
		}
	})
}
