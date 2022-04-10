package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"uahSalaryBot/external/db"
	"uahSalaryBot/external/telegram"
	"uahSalaryBot/service"
	"uahSalaryBot/usecase"

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
	bot, err := telegram.NewTgBot(token)
	if err != nil {
		log.Panic("Token is not valid")
	}

	log.Printf("Authorized on account %s", os.Getenv("BOT_NAME"))

	grace := make(chan os.Signal, 1)
	signal.Notify(grace, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		logrus.Errorf("system call: %+v", <-grace)
		cancel()
	}()

	commands := setDependencies(bot)

	server := service.NewManagerServer(bot, commands)
	server.Run(ctx)
}

func setDependencies(bot *telegram.TelegramBot) *service.Commands {
	dbClient := db.NewDbClient()

	repositories := service.NewRepositories(dbClient)
	clients := usecase.NewClients(bot, repositories)
	useCases := service.NewUseCases(clients)
	commands := service.NewCommands(useCases)

	return commands
}
