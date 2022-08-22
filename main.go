package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	token := goDotEnvVariable("TELEGRAM_BOT_API_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Printf("panic")
		panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "/test" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "awesome, bot is working")
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
		if update.Message.Text == "/send" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "send")
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
