package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	apiKey := os.Getenv("BOT_API_KEY")

	//TODO: provide context walking
	//ctx, cancel := context.WithCancel(context.Background())

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(fmt.Printf("Api key is incorrect: %s", err.Error()))
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Русский военный корабль иди на хуй")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
