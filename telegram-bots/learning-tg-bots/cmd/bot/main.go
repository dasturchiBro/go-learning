package main 

import (
	"log"
	"os"
	"xlsxbot/db"
	"xlsxbot/app"
	"xlsxbot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func AppInit() {
	var err error
	app.DB, err = db.Connect()
	if err != nil {
		log.Fatal("couldn't connect to db: ", err)
	}
}

func main() {
	AppInit()
	defer app.DB.Close()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_API"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	log.Print("BOT has started...")
	for update := range updates {
		if update.Message != nil {
			exists, err := db.UserExistsByUserID(update.Message.Chat.ID)
			if err != nil {
				log.Fatal(err)
			}
			if !exists {
				if update.Message.Contact != nil {
					if update.Message.Contact.UserID == update.Message.Chat.ID {
						err, err2 := handlers.RegisterUser(update.Message.Chat.ID, bot)
						if err != nil {
							log.Print(err)
						}
						if err2 != nil {
							log.Print(err2)
						}
					}
				} else {
					err := handlers.SendPhoneNumberButton(update.Message.Chat.ID, bot)
					if err != nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. Please try again.")
						bot.Send(&msg)
						log.Printf("An error occured: %v", err)
					}
				}
			} else {
				db.SetStageByUserID(update.Message.Chat.ID, "main")
				stage, err := db.GetStageByUserID(update.Message.Chat.ID)
				if err != nil {
					log.Print(err)
				}
				if update.Message.Text == "📚 Classes" && stage == "main" {
					err := handlers.ShowClasses(update.Message.Chat.ID, bot)
					if err != nil {
						log.Printf("An error occured in Classes Query Method: %v", err)
					}
				} 
			}
		} else if update.CallbackQuery != nil {
			log.Print(update.CallbackQuery)
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			_, err := bot.Request(callback)
			if err != nil {
				log.Print(err)
			}

			if update.CallBack.Query.Data == "Add Class CallBack" {
				// Write a handlers function to add a new class.
			} else if update.CallBack.Data == "Remove Class CallBack" {
				// Write a handlers function to remove a class with its students.
			}
		}
	}
}