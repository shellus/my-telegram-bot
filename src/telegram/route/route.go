package route

import (
	"github.com/shellus/my-telegram-bot/src/telegram/router"
	"gopkg.in/telegram-bot-api.v4"
	"github.com/astaxie/beego/logs"
	"bytes"
	"text/template"
	"fmt"
)


var Commands = router.NewRouter()
var Texts = router.NewRouter()

func Command(pattern string, handle func(update tgbotapi.Update)) (*router.Route) {
	pattern = fmt.Sprintf(`^%s(.*?)$`, pattern)
	return Commands.Add(pattern, handle)
}
func Text(pattern string, handle func(update tgbotapi.Update)) (*router.Route) {
	return Texts.Add(pattern, handle)
}

//
func Dispatch(update tgbotapi.Update) {
	logs.Info("[%s] %s %d", update.Message.From.String(), update.Message.Text, update.Message.Chat.ID)

	uri := update.Message.Text

	var (
		ro *router.Route
		err error
	)
	if uri[0:1] == "/" {
		ro, err = Commands.Dispatch(uri)
	} else {
		ro, err = Texts.Dispatch(uri)
	}

	if err != nil {
		logs.Error(err)
	}else {
		ro.Handle(update)
	}
}

func GetCommandHelpStr() (string) {
	help := bytes.NewBufferString("")

	const letter string = `{{.Pattern}} : {{.Comment}}` + "\n"
	t := template.Must(template.New("letter").Parse(letter))
	for _, r := range Commands.Routes() {
		t.Execute(help, r)
	}

	return help.String()
}