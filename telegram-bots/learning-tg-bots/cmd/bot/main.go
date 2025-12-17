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
				if update.Message.Text == "🏡 Main Menu" || update.Message.Text == "/start" {
					db.SetStageByUserID(update.Message.Chat.ID, "main")
					handlers.GoToMainMenu(update.Message.Chat.ID, bot)
				} else if update.Message.Text == "📚 Classes"{
					db.SetStageByUserID(update.Message.Chat.ID, "main")
					err := handlers.ShowClasses(update.Message.Chat.ID, bot)
					if err != nil {
						log.Printf("An error occured in Classes Query Method: %v", err)
					}
				} else if update.Message.Text == "📐 Templates" {
					db.SetStageByUserID(update.Message.Chat.ID, "main")
					err := handlers.ShowTemplates(update.Message.Chat.ID, bot)
					if err != nil {
						log.Print(err)
					}
				} else if stage == "add_class" {
					message := update.Message.Text
					parts := strings.Split(strings.TrimSpace(message), "\n")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					success := false
					if len(parts) != 2 {
						msg.Text = "Please enter the class information in the correct order: \n\t\tClass Name (e.g 11.01-E1)\n\t\tGrade in numbers (e.g. 11)"
						msg.ReplyMarkup = handlers.CancelKeyboard
					} else {
						num, err := strconv.Atoi(parts[1])
						if err != nil {
							msg.Text = "Grade should be a number: \n\t\tClass Name (e.g 11.01-E1)\n\t\tGrade (e.g. 11)\nTry again"
							msg.ReplyMarkup = handlers.CancelKeyboard
						} else {
							_, err := db.AddClass(update.Message.Chat.ID, parts[0], num)
							if err != nil {
								msg.Text = "An error has occured."
								log.Print(err)
							} else {
								msg.Text = "Class added successfully."
								success = true
							}
						}
					}
					_, err := bot.Send(&msg)
					if err != nil {
						log.Print(err)
						success = false
					}
					
					if success {
						err = handlers.ShowClasses(update.Message.Chat.ID, bot)
						if err != nil {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. Please try again later.")
							bot.Send(&msg)
						}
						_, err = db.SetStageByUserID(update.Message.Chat.ID, "main")
						if err != nil {
							msg.Text = "Something went wrong. Please try again later."
						}
					}
				} else if stage == "remove_class" {
					// START: REMOVE CLASS STAGE HANDLER
					message := update.Message.Text 
					id, err := strconv.Atoi(message)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "The class was successfully removed.")
					success := true
					if err != nil {
						success = false
						msg.Text = "ID must be an integer. Please try again."
						msg.ReplyMarkup = handlers.CancelKeyboard
					}
					r_success, err := db.RemoveClass(update.Message.Chat.ID, id)
					if err != nil {
						success = false
						msg.Text = "Something went wrong. Please try again."
						msg.ReplyMarkup = handlers.CancelKeyboard
					} else if r_success == false {
						success = false
						msg.Text = "Class with this ID doesn't exist. Please try to enter a valid ID."
						msg.ReplyMarkup = handlers.CancelKeyboard
					} 

					_, err = bot.Send(&msg)
					if err != nil {
						log.Printf("An error occured while sending the message: %v ", err)
					}
					if success {
						err := handlers.ShowClasses(update.Message.Chat.ID, bot)
						if err != nil {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "An error occured in the system.")
							bot.Send(&msg)
							log.Printf("An error occured while sending classes: %v", err)
						}
						_, err = db.SetStageByUserID(update.Message.Chat.ID, "main")
						if err != nil {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "An error occured in the system.")
							bot.Send(&msg)
							log.Printf("An error occured while setting the stage to main: %v", err)
						}
					}
				}
				// END: REMOVE CLASS STAGE HANDLER

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
				cancelKeyboard := handlers.CancelKeyboard
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "- Enter the information about the class in this order -\nClass Name\nClass Grade (in numbers)")
				msg.ReplyMarkup = cancelKeyboard
				_, err = bot.Send(&msg)
				if err != nil {
					log.Print(err)
				}
			} else if update.CallbackQuery.Data == "Remove Class Callback" {
				exists, err := db.SetStageByUserID(update.CallbackQuery.From.ID, "remove_class")
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Enter the ID of the class you want to remove:")
				if err != nil {
					log.Print(err)
					msg.Text = "Something went wrong. Please try again later."
					msg.ReplyMarkup = handlers.CancelKeyboard
				}
				if !exists {
					log.Printf("User with ID %v doesn't exist", update.CallbackQuery.From.ID)
					msg.Text = "You are not registered to this bot. Please enter /start to register to the bot."
				}
				msg.ReplyMarkup = handlers.CancelKeyboard
				_, err = bot.Send(&msg)
				if err != nil {
					log.Printf("An error occured while sending the message: %v", err)
				}
			} else if update.CallbackQuery.Data == "Return to the main menu." {
				err := handlers.GoToMainMenu(update.CallbackQuery.From.ID, bot)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}