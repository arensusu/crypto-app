package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"

	"funding-rate/controller"
	"funding-rate/database"
	"funding-rate/model"
)

var (
	tgApiToken = os.Getenv("TELEGRAM_APITOKEN")
	debug      = true
)

func loadDatabase() {
	database.Connect()
	database.DB.AutoMigrate(&model.WatchList{})
}

func main() {
	loadDatabase()

	bot, err := tgbotapi.NewBotAPI(tgApiToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		reply := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			reply.Text = "I understand /sayhi and /status."
		case "funding":
			reply.Text = controller.Funding(update.Message.Chat.ID)
		case "newfunding":
			reply.Text = controller.NewFunding(update.Message.Chat.ID, update.Message.Text)
		default:
			reply.Text = "I don't know that command"
		}

		if _, err := bot.Send(reply); err != nil {
			fmt.Println(err)
		}
	}
}
