package main

import (
	"log"
	// "os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5331249825:AAHRhfUCExDd4OEqVXcpvgvO7W4jT_E4sPw")
	if err != nil {
		log.Panic(err)
	}

	// Debugging mode
	bot.Debug = true
	//

	log.Printf("Authorised on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	var Keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Change volume"),
			tgbotapi.NewKeyboardButton("Change brightness"),
		),
	)

	// var numericInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonURL("1.com", "https://1.com"),
	// 		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
	// 		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	// 	),
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
	// 		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
	// 		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	// 	),
	// )

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = Keyboard
			case "close":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}

			switch update.Message.Command() {
			case "help":
				msg.Text = "I understand /sayhi and /status."
				msg.ReplyToMessageID = update.Message.MessageID

			case "sayhi":
				msg.Text = "HI! :-)"

			case "status":
				msg.Text = "Im ok."

			case "vol":
				continue

			case "brght":
				continue

				// case "show":
				// 	msg.ReplyMarkup = numericInlineKeyboard
				// 	// default:
				// 	// 	msg.Text = "Sorry, I don't understand you."
				// 	// 	msg.ReplyToMessageID = update.Message.MessageID
				// }

				// if _, err := bot.Send(msg); err != nil {
				// 	log.Panic(err)
				// }
				// } else if update.CallbackQuery != nil {
				// 	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				// 	if _, err := bot.Request(callback); err != nil {
				// 		panic(err)
				//	 	}

				// 	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
				// 	if _, err := bot.Send(msg); err != nil {
				// 		panic(err)
				// 	}
			}
		}
	}
}
