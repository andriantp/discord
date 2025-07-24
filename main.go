package main

import (
	"context"
	"discord/pkg/discord"
	"discord/pkg/logger"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	token := os.Getenv("token")
	if token == "" {
		logger.Level("fatal", "main", "token was required")
	}

	//for filter/send
	var channelID = map[string]string{
		"channel": "",
	}

	ctx, cancel := context.WithCancel(context.Background())
	_, err := discord.NewRepo(token, channelID)
	if err != nil {
		logger.Level("error", "main", fmt.Sprintf("discord.NewRepo:%v", err))
		cancel()
	}
	logger.Level("info", "Start", "succes connect discord")

	time.Sleep(1 * time.Second)

	log.Println("===== Run  =====")
	<-ctx.Done()
}
