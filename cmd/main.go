package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"uahSalaryBot/infrastructure"
	"uahSalaryBot/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal("Could not load .env file")
	}

	token := os.Getenv("BOT_API_KEY")
	if token == "" {
		log.Panic("Tg token is not set")
	}
	bot, err := service.NewTgBot(token)
	if err != nil {
		log.Panic("Token is not valid")
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	grace := make(chan os.Signal, 1)
	signal.Notify(grace, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		logrus.Errorf("system call: %+v", <-grace)
		cancel()
	}()

	commandManager := infrastructure.NewManager(bot)
	commandManager.Run(ctx, bot)
}
