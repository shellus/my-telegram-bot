package telegram

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"github.com/shellus/pkg/sshrsa"
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

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	//msg := tgbotapi.NewMessage(430914406, "hi 你大爷")
	//bot.Send(msg)
	//os.Exit(1)

	// 异步查询比特币
	bitcoinQueryChan := make(chan tgbotapi.Update, 10)
	go listenBitcoinQuery(bitcoinQueryChan)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s %d", update.Message.From.String(), update.Message.Text, update.Message.Chat.ID)

		var msg tgbotapi.MessageConfig

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID,
					`
					/start : 获取命令列表
					/bitcoin : 查询比特币价格`)
				bot.Send(msg)
				continue
				break
			case "bitcoin":
				bitcoinQueryChan <- update
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "等下，我正在查")
				break
			default:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("未知命令%s", update.Message.Command()))
				break
			}
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "哦，你在说什么啊？")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}

func listenBitcoinQuery(bitcoinQueryChan chan tgbotapi.Update){
	for update := range bitcoinQueryChan {
		queryBitcoinPrice(update)
	}
}
