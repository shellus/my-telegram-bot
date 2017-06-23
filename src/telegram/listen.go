package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
	route"github.com/shellus/my-telegram-bot/src/telegram/route"
	"github.com/astaxie/beego/logs"
)

func listenUpdates(){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logs.Alert("bot.GetUpdatesChan(u) err: %#v", err)
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		route.Handle(update)

	}
}