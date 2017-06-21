package telegram

import (
	route "github.com/shellus/my-telegram-bot/src/telegram/command_route"
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
)
// 异步查询比特币chan
var bitcoinQueryChan = make(chan tgbotapi.Update, 10)

func initRoutes(){
	route.Command("start", actionStart)

	// 处理比特币查询请求
	go listenBitcoinQuery(bitcoinQueryChan)
	route.Command("bitcoin", actionBitcoin)

	route.Command("default", actionDefault)
	route.Text("default", actionText)

	route.Text("哈哈", func (update tgbotapi.Update){
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "哈毛线啊")
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	})
}

func actionStart(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		`/start : 获取命令列表
		/bitcoin : 查询比特币价格`)
	bot.Send(msg)
}

func actionBitcoin(update tgbotapi.Update){
	bitcoinQueryChan <- update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "等下，我正在查")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func actionDefault(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("未知命令：{%s}", update.Message.Command()))
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func actionText(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "哦，你在说什么啊？")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}