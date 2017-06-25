package user

import (
	"github.com/astaxie/beego/orm"
	_"github.com/mattn/go-sqlite3"
	"gopkg.in/telegram-bot-api.v4"
	"encoding/json"
	"time"
)

var _ = `
CREATE TABLE "users"
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id INTEGER NOT NULL,
    chat_meta TEXT
, created_at DATETIME NULL, updated_at DATETIME NULL);
CREATE UNIQUE INDEX users_chat_id_uindex ON "users" (chat_id)
`
var DB orm.Ormer
func init(){
	orm.RegisterModel(new(User))

	//  mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//orm.RegisterDriver("sqlite3", orm.DRSqlite)

	//test.db
	//file:test.db?cache=shared&mode=memory
	//:memory:
	//file::memory:

	// todo db path
	err := orm.RegisterDataBase("default", "sqlite3", `file:C:\Users\shellus\go\src\github.com\shellus\my-telegram-bot\bin\telegram.sqlite3`)
	if err != nil {
		panic(err)
	}
	DB = orm.NewOrm()
}
type ChatMetaJson string
type User struct {
	Id int64
	Chat_id int64
	Chat_meta ChatMetaJson
	chat *tgbotapi.Chat `orm:"-"`

	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Chat()(*tgbotapi.Chat){
	chat := new(tgbotapi.Chat)
	if u.chat == nil {
		err := json.Unmarshal([]byte(u.Chat_meta), &chat)
		if err != nil {
			panic(err)
		}
		u.chat = chat
	}
	return u.chat
}