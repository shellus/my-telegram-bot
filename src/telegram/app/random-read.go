package app

import (
	"gopkg.in/telegram-bot-api.v4"
	"github.com/shellus/my-telegram-bot/src/telegram/bot"
)

func ActionRandomRead(update tgbotapi.Update){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "你看毛线看")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Bot.Send(msg)
}

//
//func getIndex(){
//	http.Get("")
//}