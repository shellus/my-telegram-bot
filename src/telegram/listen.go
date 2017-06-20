package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
	route"github.com/shellus/my-telegram-bot/src/telegram/command_route"
	"log"
)

func Listen(){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s %d", update.Message.From.String(), update.Message.Text, update.Message.Chat.ID)
		route.Handle(update)
	}
}