package main 

import (
	"log"
	"os"
	"xlsxbot/db"
	"xlsxbot/app"
	"xlsxbot/handlers"
	"xlsxbot/models"
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
				} /* END: REMOVE CLASS STAGE HANDLER */ else if strings.Contains(stage, "add students to class ") {

				 // ****START ADD STUDENTS STAGE HANDLER**** //
					parts := strings.Fields(stage)
					id, _ := strconv.Atoi(parts[len(parts) - 1])
					message := update.Message.Text
					students := strings.Split(message, "\n")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Student(s) added successfully!")
					for _, student := range students {
						var newStudent models.Student
						newStudent.Name = student
						newStudent.ClassID = id
						newStudent.UserID = update.Message.Chat.ID
						_, err := db.AddStudentToUser(newStudent)
						if err != nil {
							log.Printf("An error occured while adding student: %v", err)
							msg.Text = "Something went wrong. Please try again later."
						}
					}
					_, err := bot.Send(&msg)
					if err != nil {
						log.Printf("An error occured while sending a message: %v", err)
					}
					
					if _, err := db.SetStageByUserID(update.Message.Chat.ID, "main"); err != nil {
						log.Print(err)
					}
					
					if err := handlers.ShowClass(update.Message.Chat.ID, bot, id); err != nil {
						log.Printf("An error occured in Classes Query Method: %v", err)
					}
				
				}/* ****END ADD STUDENTS STAGE HANDLER***** */ else if strings.Contains(stage, "remove students from class ") {

				 // ****START REMOVE STUDENT STAGE HANDLER**** //

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Student removed successfully!")
					parts := strings.Fields(stage)
					classID, _ := strconv.Atoi(parts[len(parts) - 1])
					number, err := strconv.Atoi(update.Message.Text)
					if err != nil {
						msg.Text = "The student's number should be an integer. Please try again."
					}
					students, err := db.GetStudentsByClassID(classID, update.Message.Chat.ID)
					if err != nil {
						msg.Text = "Something went wrong. Please try again later."
					}
					if number > 0 && number <= len(students) {
						student := students[number - 1]
						ok, err := db.RemoveStudentByClassID(update.Message.Chat.ID, classID, student.ID)
						if err != nil {
							log.Print(err)
							msg.Text = "Something went wrong. Please try again later."
						}
						if !ok {
							msg.Text = "An error occured: Student removal wasn't successful. "
						}
					} else {
						msg.Text = "Student on this number doesn't exist."
					}
					_, err = bot.Send(&msg)
					if err != nil {
						log.Printf("An error occured while sending a message: %v", err)
					}
					
					if _, err := db.SetStageByUserID(update.Message.Chat.ID, "main"); err != nil {
						log.Print(err)
					}
					
					if err := handlers.ShowClass(update.Message.Chat.ID, bot, classID); err != nil {
						log.Printf("An error occured in Classes Query Method: %v", err)
					}
				
				}// ****END REMOVE STUDENT STAGE HANDLER***** // 


				

				

			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			_, err := bot.Request(callback)
			if err != nil {
				log.Print(err)
			}

			if update.CallbackQuery.Data == "Add Class Callback" {
				err = handlers.DeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, bot)
				if err != nil {
					log.Printf("An error occured while deleting a message: %v", err)
				}
				exists, err := db.SetStageByUserID(update.CallbackQuery.From.ID, "add_class")
				if err != nil {
					log.Print(err)
				}
				if !exists {
					log.Print("User with this id doesn't exist")
				}
				cancelKeyboard := handlers.ReturnToClassesKeyboard
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "- Enter the information about the class in this order -\nClass Name\nClass Grade (in numbers)")
				msg.ReplyMarkup = cancelKeyboard
				_, err = bot.Send(&msg)
				if err != nil {
					log.Print(err)
				}
			} else if update.CallbackQuery.Data == "Remove Class Callback" {
				err = handlers.DeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, bot)
				if err != nil {
					log.Printf("An error occured while deleting a message: %v", err)
				}
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
				msg.ReplyMarkup = handlers.ReturnToClassesKeyboard
				_, err = bot.Send(&msg)
				if err != nil {
					log.Printf("An error occured while sending the message: %v", err)
				}
			} else if update.CallbackQuery.Data == "Return to the main menu." {
				err := handlers.GoToMainMenu(update.CallbackQuery.From.ID, bot)
				if err != nil {
					log.Print(err)
				}
				
				err = handlers.DeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, bot)
				if err != nil {
					log.Printf("An error occured while deleting a message: %v", err)
				}
			} else if update.CallbackQuery.Data == "Return to classes." {
				err = handlers.DeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, bot)
				if err != nil {
					log.Printf("An error occured while deleting a message: %v", err)
				}
				db.SetStageByUserID(update.CallbackQuery.From.ID, "main")
				err := handlers.ShowClasses(update.CallbackQuery.From.ID, bot)
				if err != nil {
					log.Printf("An error occured in Classes Query Method: %v", err)
				}
				
			} else if strings.Contains(update.CallbackQuery.Data, "Manage class with ID ") {
				err = handlers.DeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, bot)
				if err != nil {
					log.Printf("An error occured while deleting a message: %v", err)
				}
				parts := strings.Fields(update.CallbackQuery.Data)
				id := parts[len(parts) - 1]
				classID, _ := strconv.Atoi(id)
				err := handlers.ShowClass(update.CallbackQuery.From.ID, bot, classID)
				if err != nil {
					log.Printf("An error occured in Class Query Method: %v", err)
				}
			} else if strings.Contains(update.CallbackQuery.Data, "Add Students to class ") {
				err := handlers.DeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, bot)
				if err != nil {
					log.Printf("An error occured while deleting a message: %v", err)
				}
				parts := strings.Fields(update.CallbackQuery.Data)
				id := parts[len(parts)-1]
				// classID, _ := strconv.Atoi(id)
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Enter each student’s name on a new line:")
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Return back.", "Manage class with ID " + id),
					),
				)
				_, err = db.SetStageByUserID(update.CallbackQuery.From.ID, "add students to class " + id)
				if err != nil {
					log.Printf("Something went from while setting the stage: %v", err)
				}
				if _, err := bot.Send(&msg); err != nil {
					log.Printf("Something went wrong while sending message: %v", err)
				}
			} else if strings.Contains(update.CallbackQuery.Data, "Remove Students from class ") {
				parts := strings.Fields(update.CallbackQuery.Data)
				id := parts[len(parts)-1]
				// classID, _ := strconv.Atoi(id)
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Enter the student’s number you want to delete:")
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Return back.", "Manage class with ID " + id),
					),
				)
				_, err = db.SetStageByUserID(update.CallbackQuery.From.ID, "remove students from class " + id)
				if err != nil {
					log.Printf("Something went from while setting the stage: %v", err)
				}
				if _, err := bot.Send(&msg); err != nil {
					log.Printf("Something went wrong while sending message: %v", err)
				}
			} else if strings.Contains(update.CallbackQuery.Data, "Delete Class ") {
				err := handlers.DeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, bot)
				if err != nil {
					log.Printf("An error occured while deleting a message: %v", err)
				}
				parts := strings.Fields(update.CallbackQuery.Data)
				id := parts[len(parts)-1]
				classID, _ := strconv.Atoi(id)
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "The class was successfully removed.")
				success := true
				r_success, err := db.RemoveClass(update.CallbackQuery.From.ID, classID)
				if err != nil {
					success = false
					msg.Text = "Something went wrong. Please try again."

				} else if r_success == false {
					success = false
					msg.Text = "Class with this ID doesn't exist. Please try to enter a valid ID."
				} 

				_, err = bot.Send(&msg)
				if success == false {
					err := handlers.ShowClass(update.CallbackQuery.From.ID, bot, classID)
					if err != nil {
						log.Printf("An error occured in Class Query Method: %v", err)
					}
				}
				if err != nil {
					log.Printf("An error occured while sending the message: %v ", err)
				}
				if success {
					err := handlers.ShowClasses(update.CallbackQuery.From.ID, bot)
					if err != nil {
						msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "An error occured in the system.")
						bot.Send(&msg)
						log.Printf("An error occured while sending classes: %v", err)
					}
					_, err = db.SetStageByUserID(update.CallbackQuery.From.ID, "main")
					if err != nil {
						msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "An error occured in the system.")
						bot.Send(&msg)
						log.Printf("An error occured while setting the stage to main: %v", err)
					}
				}
			}
		}
	}
}