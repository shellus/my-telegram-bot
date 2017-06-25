package user

import (
	"gopkg.in/telegram-bot-api.v4"
	"github.com/jmoiron/sqlx"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	_"github.com/mattn/go-sqlite3"
)
var db *sqlx.DB

var _ = `
CREATE TABLE "users"
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id INTEGER NOT NULL,
    chat_meta TEXT
);
CREATE UNIQUE INDEX users_chat_id_uindex ON "users" (chat_id)
`

func init(){
	var err error

	//test.db
	//file:test.db?cache=shared&mode=memory
	//:memory:
	//file::memory:

	// todo db path
	db, err = sqlx.Connect("sqlite3", `file:bin/telegram.sqlite3`)
	if err != nil {
		panic(err)
	}
}
type User struct {
	Id int64
	Chat_id int64
	TgChat *tgbotapi.Chat
}

func Create(u *User)(uR *User){
	c, err :=  json.Marshal(u.TgChat)
	if err != nil {
		panic(c)
	}
	u.Id, err = db.MustExec("INSERT INTO users (chat_id, chat_meta) VALUES ($1, $2);", u.Chat_id, string(c)).LastInsertId()
	if err != nil {
		panic(c)
	}
	uR = u
	return
}
func FindChatId(id int64)(u *User, err error) {
	var chatJs = ""
	u = &User{TgChat:&tgbotapi.Chat{}}
	err = db.QueryRow("SELECT * FROM users WHERE chat_id=$1", id).Scan(&u.Id, &u.Chat_id, &chatJs)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(chatJs), u.TgChat)
	if err != nil {
		logs.Error(chatJs)
		panic(err)
	}
	return
}

func (u *User) Delete()(bool){
	rows, err := db.MustExec("DELETE FROM users WHERE id = $1;", u.Id).RowsAffected()
	if err != nil {
		panic(err)
	}
	return rows != 1
}