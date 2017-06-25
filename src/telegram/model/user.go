package model

import (
	"gopkg.in/telegram-bot-api.v4"
	"github.com/jmoiron/sqlx"
	"encoding/json"
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

func Create(u *User)(bool){
	c, err :=  json.Marshal(u.TgChat)
	if err != nil {
		panic(c)
	}
	u.Id, err = db.MustExec("INSERT INTO users (chat_id, chat_meta) VALUES ($1, $2);", u.Chat_id, string(c)).LastInsertId()
	if err != nil {
		panic(c)
	}
	return true
}

func (u *User) Delete()(bool){
	rows, err := db.MustExec("DELETE FROM users WHERE id = $1;", u.Id).RowsAffected()
	if err != nil {
		panic(err)
	}
	return rows != 1
}