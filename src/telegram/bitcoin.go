package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"net/http"
	"github.com/shellus/jason"
)

func queryBitcoinPrice(update tgbotapi.Update){
	var msg tgbotapi.MessageConfig
	defer func(){
		if r := recover(); r != nil{
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "查询出错了")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
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
	ads := j.MustGetObject("data").MustGetObjectArray("ad_list")
	price_buy_high := ads[0].MustGetObject("data").MustGetString("temp_price")

	resp, err = http.Get("https://localbitcoins.com/sell-bitcoins-online/cn/china/zhi-fu-bao/.json")
	if err != nil {
		panic(err)
	}
	j, err = jason.NewObjectFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	ads = j.MustGetObject("data").MustGetObjectArray("ad_list")
	price_sell_low := ads[0].MustGetObject("data").MustGetString("temp_price")

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("localbitcoins.com\n最低买入: ￥%s\n最高出售: ￥%s", price_sell_low, price_buy_high))
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}