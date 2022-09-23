package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func VolumeHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.Text = "Enter volume level:"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	answ := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	level := answ.Text

}
