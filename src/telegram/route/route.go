package route

import (
	"github.com/shellus/my-telegram-bot/src/telegram/router"
	"gopkg.in/telegram-bot-api.v4"
	"github.com/astaxie/beego/logs"
	"bytes"
	"text/template"
)

var Commands = router.NewRouter()
var Texts = router.NewRouter()

func Handle(update tgbotapi.Update){
	logs.Info("[%s] %s %d", update.Message.From.String(), update.Message.Text, update.Message.Chat.ID)

	if update.Message.IsCommand() {
		Commands.Handle(update)
	}else {
		Texts.Handle(update)
	}
}

func GetCommandHelpStr()(string){
	help := bytes.NewBufferString("")

	const letter string = `{{.Pattern}} : {{.Comment}}` + "\n"
	t := template.Must(template.New("letter").Parse(letter))
	for _, r := range Commands.Routes(){
		t.Execute(help,r)
	}

	return help.String()
}