package telegram

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
)

var bot *tgbotapi.BotAPI



func Main() {
	var err error

	bot, err = tgbotapi.NewBotAPI(gettoken())

	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	initRoutes()
	listenUpdates()
}

func Init(){

}


