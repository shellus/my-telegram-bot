package telegram

import (
	route "github.com/shellus/my-telegram-bot/src/telegram/route"
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
)
// 异步查询比特币chan
var bitcoinQueryChan = make(chan tgbotapi.Update, 10)

func initRoutes(){
	route.Commands.Add("/start", actionStart).SetComment("获取命令列表")

	// 处理比特币查询请求
	go listenBitcoinQuery(bitcoinQueryChan)
	route.Commands.Add("/bitcoin", actionBitcoin).SetComment("查询比特币价格")

	route.Commands.Default(actionDefault)


	route.Texts.Add("哈哈", actionHAHA)

	route.Texts.Add("default", actionText)
}

func actionStart(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, route.GetCommandHelpStr())
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

func actionHAHA (update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "哈毛线啊")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}
func actionText(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "哦，你在说什么啊？")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}