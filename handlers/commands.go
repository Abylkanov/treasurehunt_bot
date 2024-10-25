package handlers

import (
	"log"
	//"database/sql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) string {
	message := update.Message
	// Логирование входящих сообщений
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	// Ответ на команду /start
	if message.IsCommand() {
		switch message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Введите год выпуска серии Hot Wheels \n ⬇️⬇️⬇️.")
			bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(message.Chat.ID, "Команда не распознана.")
			bot.Send(msg)
		}
	}
	return "year"
}
