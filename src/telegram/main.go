package telegram

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
)

var bot *tgbotapi.BotAPI

// 异步查询比特币chan
var bitcoinQueryChan = make(chan tgbotapi.Update, 10)

func Main() {
	var err error

	bot, err = tgbotapi.NewBotAPI(gettoken())

	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// 处理比特币查询请求
	go listenBitcoinQuery(bitcoinQueryChan)

	Listen()
}


