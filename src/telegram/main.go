package telegram

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"github.com/astaxie/beego/logs"
	"github.com/shellus/my-telegram-bot/src/telegram/bot"
)


func Main() {
	var err error

	bot.Bot, err = tgbotapi.NewBotAPI(gettoken())

	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	logs.Notice("Authorized on account %s", bot.Bot.Self.UserName)

	initRoutes()
	listenUpdates()
}

