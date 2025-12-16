package main 

import (
	"log"
	"os"
	"xlsxbot/db"
	"xlsxbot/app"
	"xlsxbot/handlers"
	"strings"
	"strconv"
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
				stage, err := db.GetStageByUserID(update.Message.Chat.ID)
				if err != nil {
					log.Print(err)
				}
				if update.Message.Text == "🏡 Main Menu" {
					db.SetStageByUserID(update.Message.Chat.ID, "main")
					handlers.GoToMainMenu(update.Message.Chat.ID, bot)
				} else if update.Message.Text == "📚 Classes" && stage == "main" {

					err := handlers.ShowClasses(update.Message.Chat.ID, bot)
					if err != nil {
						log.Printf("An error occured in Classes Query Method: %v", err)
					}
				} else if stage == "add_class" {
					message := update.Message.Text
					parts := strings.Split(strings.TrimSpace(message), "\n")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					if len(parts) != 2 {
						msg.Text = "Please enter the class information in the correct order: \n\t\tClass Name (e.g 11.01-E1)\n\t\tGrade in numbers (e.g. 11)"
					} else {
						num, err := strconv.Atoi(parts[1])
						if err != nil {
							msg.Text = "Grade should be a number: \n\t\tClass Name (e.g 11.01-E1)\n\t\tGrade (e.g. 11)"
						} else {
							_, err := db.AddClass(update.Message.Chat.ID, parts[0], num)
							if err != nil {
								msg.Text = "An error has occured."
								log.Print(err)
							} else {
								msg.Text = "Class added successfully."
								_, err = db.SetStageByUserID(update.Message.Chat.ID, "main")
								if err != nil {
									msg.Text = "Something went wrong. Please try again later."
								}
							}
						}
					}
					_, err := bot.Send(&msg)
					if err != nil {
						log.Print(err)
					}
					
					err = handlers.ShowClasses(update.Message.Chat.ID, bot)
					if err != nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. Please try again later.")
						bot.Send(&msg)
					}
				}
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			_, err := bot.Request(callback)
			if err != nil {
				log.Print(err)
			}

			if update.CallbackQuery.Data == "Add Class Callback" {
				exists, err := db.SetStageByUserID(update.CallbackQuery.From.ID, "add_class")
				if err != nil {
					log.Print(err)
				}
				if !exists {
					log.Print("User with this id doesn't exist")
				}
				cancelKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Return back", "Return to the main menu.")),
				)
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "- Enter the information about the class in this order -\nClass Name\nClass Grade (in numbers)")
				msg.ReplyMarkup = cancelKeyboard
				_, err = bot.Send(&msg)
				if err != nil {
					log.Print(err)
				}
			} else if update.CallbackQuery.Data == "Remove Class Callback" {
				// Write a handlers function to remove a class with its students.
			} else if update.CallbackQuery.Data == "Return to the main menu." {
				err := handlers.GoToMainMenu(update.CallbackQuery.From.ID, bot)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}