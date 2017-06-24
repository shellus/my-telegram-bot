package app

import (
	"gopkg.in/telegram-bot-api.v4"
	"github.com/shellus/my-telegram-bot/src/telegram/route"
	"github.com/shellus/my-telegram-bot/src/telegram/bot"
	"fmt"
)

func ActionStart(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, route.GetCommandHelpStr())
	bot.Bot.Send(msg)
}

func ActionBitcoin(update tgbotapi.Update){
	BitcoinQueryChan <- update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "等下，我正在查")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Bot.Send(msg)
}

func ActionDefault(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("未知命令：{%s}", update.Message.Command()))
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Bot.Send(msg)
}

func ActionHAHA (update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "哈毛线啊")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Bot.Send(msg)
}
func ActionText(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "哦，你在说什么啊？")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Bot.Send(msg)
}
