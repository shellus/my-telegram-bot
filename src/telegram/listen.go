package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
	"github.com/shellus/my-telegram-bot/src/telegram/route"
	"github.com/astaxie/beego/logs"
	"github.com/shellus/my-telegram-bot/src/telegram/bot"
	"github.com/shellus/my-telegram-bot/src/telegram/model/user"
	"fmt"
)

func listenUpdates(){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.Bot.GetUpdatesChan(u)
	if err != nil {
		logs.Alert("bot.GetUpdatesChan(u) err: %#v", err)
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		checkUser(update.Message.Chat)
		route.Dispatch(update)
	}
}

func checkUser(chat *tgbotapi.Chat){

	u := &user.User{Chat_id:chat.ID}
	_, _, err := user.DB.ReadOrCreate(u, "Chat_id")
	if err != nil {
		panic(err)
	}
	user.DB.Update(u, "UpdatedAt")
	fmt.Println(u.Id)
}