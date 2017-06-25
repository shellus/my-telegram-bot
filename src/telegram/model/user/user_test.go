package user

import (
	"testing"
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	_ "github.com/mattn/go-sqlite3"
)

func TestCreate(t *testing.T) {
	js := `{"id":249603346,"type":"private","title":"","username":"","first_name":"shellus","last_name":"","all_members_are_administrators":false}`
	tgChat := &tgbotapi.Chat{}
	err := json.Unmarshal([]byte(js), tgChat)
	if err != nil {
		panic(err)
	}
	user := &User{Id:1, Chat_id:249603346, TgChat:tgChat}
	if Create(user) {
		user.Delete()
	}
}