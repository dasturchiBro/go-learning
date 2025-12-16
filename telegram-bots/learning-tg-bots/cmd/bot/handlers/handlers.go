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

	classesKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Add Class", "Add Class Callback"),
			tgbotapi.NewInlineKeyboardButtonData("Remove Class", "Remove Class Callback"),
		),
	)
	
	classes, err := db.GetClassesByUserID(chatID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, "-- Your classes -- \n\n")
	if len(classes) == 0 {
		msg.Text += "\t\tYou don't have classes."
	} else {
		for i, class := range classes {
			msg.Text += "\t\t" + strconv.Itoa(i+1) + ") Name: " + class.Name + " - Grade: " + strconv.Itoa(class.Grade) + " - ID: " + strconv.Itoa(class.ID) + "\n"
		}
	}
	msg.ReplyMarkup = classesKeyboard

	_, err = bot.Send(&msg)
	return err
}

func AddClass(chatID int64, bot *tgbotapi.BotAPI) (error) {
	return nil
}

var CancelKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Return back", "Return to the main menu.")),
)