package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetupUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	return updates
}
