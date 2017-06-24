package telegram

import (
	"github.com/shellus/my-telegram-bot/src/telegram/route"
	"github.com/shellus/my-telegram-bot/src/telegram/app"
)
// 异步查询比特币chan


func initRoutes(){
	route.Command("/start", app.ActionStart).SetComment("获取命令列表")

	// 处理比特币查询请求
	go app.ListenBitcoinQuery(app.BitcoinQueryChan)
	route.Command("/bitcoin", app.ActionBitcoin).SetComment("查询比特币价格")

	route.Command("/random_read", app.ActionRandomRead).SetComment("随机看书")

	route.CommandDefault(app.ActionDefault)


	route.Text("哈哈", app.ActionHAHA)

	route.TextDefault(app.ActionText)
}