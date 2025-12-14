package main 

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("GET WEATHER"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("HELP"),
		tgbotapi.NewKeyboardButton("ABOUT"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_API"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.Command() == "start" {
			msg.Text = "Welcome to a weather bot."
			msg.ReplyMarkup = startKeyboard
		} else {
			msg.Text = update.Message.Text
		}
		bot.Send(msg)
	}
}