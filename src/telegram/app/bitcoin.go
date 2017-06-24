package app

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"net/http"
	"github.com/antonholmquist/jason"
	"github.com/shellus/my-telegram-bot/src/telegram/bot"
)

var BitcoinQueryChan = make(chan tgbotapi.Update, 10)

func ListenBitcoinQuery(bitcoinQueryChan chan tgbotapi.Update){
	for update := range bitcoinQueryChan {
		queryBitcoinPrice(update)
	}
}


func queryBitcoinPrice(update tgbotapi.Update){
	var msg tgbotapi.MessageConfig
	defer func(){
		if r := recover(); r != nil{
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "查询出错了")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Bot.Send(msg)
		}
	}()

	resp, err := http.Get("https://localbitcoins.com/buy-bitcoins-online/cn/china/zhi-fu-bao/.json")
	if err != nil {
		panic(err)
	}
	j, err := jason.NewObjectFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	ads ,_ := j.GetObjectArray("data", "ad_list")
	price_buy_high, _ := ads[0].GetString("data", "temp_price")

	resp, err = http.Get("https://localbitcoins.com/sell-bitcoins-online/cn/china/zhi-fu-bao/.json")
	if err != nil {
		panic(err)
	}
	j, err = jason.NewObjectFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	ads ,_ = j.GetObjectArray("data", "ad_list")
	price_sell_low, _ := ads[0].GetString("data", "temp_price")

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("localbitcoins.com\n最低买入: ￥%s\n最高出售: ￥%s", price_sell_low, price_buy_high))
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Bot.Send(msg)
}