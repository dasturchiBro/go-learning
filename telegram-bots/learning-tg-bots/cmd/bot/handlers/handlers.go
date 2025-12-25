package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"xlsxbot/db"
	"xlsxbot/models"
	"strconv"
	"encoding/json"
	"strings"
)

func SendPhoneNumberButton(chatID int64, bot *tgbotapi.BotAPI) error {
	phoneButton := tgbotapi.KeyboardButton{
		Text: "Send Phone Number",
		RequestContact: true,
	}

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(phoneButton),
	)

	msg := tgbotapi.NewMessage(chatID, "Please share your phone number to register to the bot.")
	msg.ReplyMarkup = keyboard

	_, err := bot.Send(&msg)
	return err
}

func GoToMainMenu(chatID int64, bot *tgbotapi.BotAPI) error {
	db.SetStageByUserID(chatID, "main")
	msg := tgbotapi.NewMessage(chatID, "*Main Menu \n\n- Choose one of the options -\n\t\t1) Classes - to see your classes\n\t\t2) Templates - to see Excel templates\n\t\t3) Help - to see instructions.")
	mainMenuKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("📚 Classes")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("📐 Templates")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("❓ Help")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("🏡 Main Menu")),
	)
	msg.ReplyMarkup = mainMenuKeyboard

	_, err := bot.Send(&msg)
	return err
}

func RegisterUser(chatID int64, bot *tgbotapi.BotAPI) (error, error) {
	err := GoToMainMenu(chatID, bot)
	_, err2 := db.InsertUser(chatID)
	return err, err2
}

func ShowClasses(chatID int64, bot *tgbotapi.BotAPI) (error) {

	classesKeyboard_main := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Add Class", "Add Class Callback"),
			tgbotapi.NewInlineKeyboardButtonData("Remove Class", "Remove Class Callback"),
		)

	var classesKeyboard [][]tgbotapi.InlineKeyboardButton
	classesKeyboard = append(classesKeyboard, classesKeyboard_main)
	
	classes, err := db.GetClassesByUserID(chatID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, "-- Your classes -- \n\n")
	if len(classes) == 0 {
		msg.Text += "\t\tYou don't have classes."
	} else {
		for i, class := range classes {
			grade := strconv.Itoa(class.Grade)
			classid := strconv.Itoa(class.ID)
			msg.Text += "\t\t" + strconv.Itoa(i+1) + ") Name: " + class.Name + " - Grade: " + grade + " - ID: " + classid + "\n"
			button := tgbotapi.NewInlineKeyboardButtonData("Manage " + class.Name + " - ID: " + classid, "Manage class with ID " + classid)
			row := tgbotapi.NewInlineKeyboardRow(button)
			classesKeyboard = append(classesKeyboard, row)
		}
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(classesKeyboard...)

	_, err = bot.Send(&msg)
	return err
}

func ShowClass(chatID int64, bot *tgbotapi.BotAPI, id int) (error) {
	classesKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Add Students", "Add Students to class " + strconv.Itoa(id)),
			tgbotapi.NewInlineKeyboardButtonData("Remove Students", "Remove Students from class " + strconv.Itoa(id)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Delete Class", "Delete Class " + strconv.Itoa(id)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Return back", "Return to classes."),
		),
	)
	class, err := db.GetClassByUserID(chatID, id)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatID, "-- Your class -- \n\n")
	grade := strconv.Itoa(class.Grade)
	classid := strconv.Itoa(class.ID)
	students, err := db.GetStudentsByClassID(class.ID, chatID)
	if err != nil {
		return err
	}
	studentsShow := ""
	if len(students) != 0 {
		studentsShow += "\n\n-- Students --\n"
		for i, student := range students {
			studentsShow += "\t\t(" + strconv.Itoa(i+1) + ") " + student.Name + "\n"
		}
	}
	msg.Text += "\t\t" + "Name: " + class.Name + "\n\t\tGrade: " + grade + "\n\t\tID: " + classid + "\n\t\tNumber of students in the class: " + strconv.Itoa(len(students)) + studentsShow
	msg.ReplyMarkup =  classesKeyboard

	_, err = bot.Send(&msg)
	return err
}



func ShowTemplates(chatID int64, bot *tgbotapi.BotAPI) error {
	templatesKeyboard_main := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Add Template", "Add Template Callback"),)
	var templatesKeyboard [][]tgbotapi.InlineKeyboardButton
	templatesKeyboard = append(templatesKeyboard, templatesKeyboard_main)
	


	templates, err := db.GetTemplatesByUserID(chatID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, "-- Your templates -- \n\n")
	if len(templates) == 0 {
		msg.Text += "\t\tYou don't have templates."
	} else {
		for i, template := range templates {
			msg.Text += "\t\t" + strconv.Itoa(i+1) + ") Name: " + template.Name + " - ID: " + strconv.Itoa(template.ID) + "\n"
			button := tgbotapi.NewInlineKeyboardButtonData("Manage " + template.Name + " - ID: " + strconv.Itoa(template.ID), "Manage template with ID " + strconv.Itoa(template.ID))
			row := tgbotapi.NewInlineKeyboardRow(button)
			templatesKeyboard = append(templatesKeyboard, row)
		}
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(templatesKeyboard...)

	_, err = bot.Send(&msg)
	return err
}

func ShowTemplate(chatID int64, bot *tgbotapi.BotAPI, id int) (error) {
	templatesKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Use Template", "Use template " + strconv.Itoa(id)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Delete Template", "Delete Template " + strconv.Itoa(id)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Return back", "Return to templates."),
		),
	)
	template, err := db.GetTemplateByUserID(chatID, id)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatID, "-- Your template -- \n\n")
	templateid := strconv.Itoa(template.ID)
	var criteria []string
	err = json.Unmarshal(template.Criteria, &criteria)
	if err != nil {
		return err
	}
	criteriaShow := "\n\tCriteria:\n"
	for _, c := range criteria {
		criteriaShow += "\t\t" + c + "\n"
	}
	msg.Text += "\t" + "Name: " + template.Name + "\n\tGrade: " + template.Header[0] + "\n\tQuarter: " + template.Header[1] + "\n\tExam Type: " + template.Header[2] + "\n\tID: " + templateid + criteriaShow
	msg.ReplyMarkup =  templatesKeyboard

	_, err = bot.Send(&msg)
	return err
}


func DeleteMessage(chatID int64, messageID int, bot *tgbotapi.BotAPI) error {
	deleteConfig := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := bot.Request(deleteConfig)
	return err
}


var CancelKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Return back", "Return to the main menu.")),
)

var ReturnToClassesKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Return back", "Return to classes.")),
)

var ReturnToTemplatesKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Return back", "Return to templates.")),
)


func UseTemplate(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	parts := strings.Fields(update.CallbackQuery.Data)
	templateID_str := parts[len(parts) - 1]
	// templateID, _ := strconv.Atoi(parts[len(parts) - 1])
	classes, err := db.GetClassesByUserID(update.CallbackQuery.From.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "-- Enter the class ID you want to use the template with -- \n\n")
	if len(classes) == 0 {
		msg.Text += "\t\tYou don't have classes. Go to Classes to add a new class."
	} else {
		for i, class := range classes {
			grade := strconv.Itoa(class.Grade)
			classid := strconv.Itoa(class.ID)
			msg.Text += "\t\t" + strconv.Itoa(i+1) + ") Name: " + class.Name + " - Grade: " + grade + " - ID: " + classid + "\n"
		}
		_, err := db.SetStageByUserID(update.CallbackQuery.From.ID, "Template_Usage_Enter_Class_ID " + templateID_str)
		if err != nil {
			return err
		}
	}	
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Return back.", "Manage template with ID " + templateID_str),
		),
	)
	_, err = bot.Send(&msg)
	return err
}

func UseTemplateClassIDHandler(stage string, update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	chatID := update.Message.Chat.ID
	parts := strings.Fields(stage)
	templateID_str := parts[len(parts) - 1]
	templateID, _ := strconv.Atoi(templateID_str)
	classID, err := strconv.Atoi(update.Message.Text)
	var mainErr error 
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if err != nil {
		msg.Text = "Please enter a correct class ID. Try again."
	} else {
		class, err := db.GetClassByUserID(chatID, classID)
		if err != nil {
			msg.Text = "Something went wrong. Please try again."
			mainErr = err
		} else if class == (models.Class{})  {
			msg.Text = "Class with ID " + update.Message.Text + " doesn't exist. Please enter a valid ID."
		} else {
			students, err := db.GetStudentsByClassID(class.ID, chatID)
			if err != nil {
				mainErr = err
			}
			if len(students) == 0 {
				msg.Text = "There are no students in the class. \nGo to Classes -> Class with ID " + update.Message.Text + " -> Add Students"
			} else {
				db.SetStageByUserID(update.Message.Chat.ID, "AddScoresToStudentTemplate_"+templateID_str+"_"+update.Message.Text+"_0_"+strconv.Itoa(len(students)))				
				template, err := db.GetTemplateByUserID(chatID, templateID)
				if err != nil {
					mainErr = err
				}
				var criteria []string
				err = json.Unmarshal(template.Criteria, &criteria)
				if err != nil {
					mainErr = err
				}
				criteriaShow := "\n\tCriteria:\n"
				for _, c := range criteria {
					criteriaShow += "\t\t" + c + "\n"
				}
				criteriaShow += "\n\nWhat scores do you want to add to the student " + students[0].Name + "? (Note that the scores should be entered in the same way as the criteria)"
				msg.Text = criteriaShow
			}
		}
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Return back.", "Manage template with ID " + templateID_str),
		),
	)
	_, err = bot.Send(&msg)
	if err != nil {
		mainErr = err	
	}
	return mainErr
}


func AddScoresToStudentTemplateHandler(stage string, update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	parts := strings.Split(stage, "_")
	templateID := parts[1]
	classID := parts[2]
	studentNumber, _ := strconv.Atoi(parts[3])
	studentsLen, _ := strconv.Atoi(parts[4])	
	templateID_int, _ := strconv.Atoi(templateID)
	classID_int, _ := strconv.Atoi(classID)
	chatID := update.Message.Chat.ID
	var mainErr error
	msg := tgbotapi.NewMessage(chatID, "")
	class, err := db.GetClassByUserID(chatID, classID_int)
	if err != nil {
		msg.Text = "Something went wrong. Please try again."
		mainErr = err
	} else if class == (models.Class{})  {
		msg.Text = "Class with ID " + update.Message.Text + " doesn't exist."
	} else {
		students, err := db.GetStudentsByClassID(class.ID, chatID)
		if err != nil {
			mainErr = err
		}
		if len(students) == 0 {
			msg.Text = "There are no students in the class. \nGo to Classes -> Class with ID " + update.Message.Text + " -> Add Students"
		} else {
			if studentNumber + 1 < studentsLen {
				db.SetStageByUserID(update.Message.Chat.ID, "AddScoresToStudentTemplate_"+templateID+"_"+update.Message.Text+"_"+strconv.Itoa(studentNumber+1)+"_"+strconv.Itoa(len(students)))				
				template, err := db.GetTemplateByUserID(chatID, templateID_int)
				if err != nil {
					mainErr = err
				}

				// ===TASK FOR 25.12.2025===
				// CHECK FOR THE CORRECT WAY OF ENTERING CRITERIA AND FOR ENTERING NUMBER not STRINGS
				// SAVE POINTS TO STUDENTS

				var criteria []string
				err = json.Unmarshal(template.Criteria, &criteria)
				if err != nil {
					mainErr = err
				}
				criteriaShow := "\n\tCriteria:\n"
				for _, c := range criteria {
					criteriaShow += "\t\t" + c + "\n"
				}
				criteriaShow += "\n\nWhat scores do you want to add to the student " + students[studentNumber+1].Name + "? (Note that the scores should be entered in the same way as the criteria)"
				msg.Text = criteriaShow

			}
		}
	}
	if _, err := bot.Send(&msg); err != nil {
		mainErr = err
	}
	return mainErr
}