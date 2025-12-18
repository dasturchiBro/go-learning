package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"xlsxbot/db"
	"strconv"
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

func AddClass(chatID int64, bot *tgbotapi.BotAPI) (error) {
	return nil
}

var CancelKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Return back", "Return to the main menu.")),
)

func ShowTemplates(chatID int64, bot *tgbotapi.BotAPI) error {
	templatesKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Add Template", "Add Template Callback"),
			tgbotapi.NewInlineKeyboardButtonData("Remove Template", "Remove Template Callback"),
		),
	)

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
		}
	}
	msg.ReplyMarkup = templatesKeyboard

	_, err = bot.Send(&msg)
	return err
}


func DeleteMessage(chatID int64, messageID int, bot *tgbotapi.BotAPI) error {
	deleteConfig := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := bot.Request(deleteConfig)
	return err
}