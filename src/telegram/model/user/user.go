package user

import (
	_"github.com/mattn/go-sqlite3"
	"gopkg.in/telegram-bot-api.v4"
	"time"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine
func init(){
	//  mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//orm.RegisterDriver("sqlite3", orm.DRSqlite)

	//test.db
	//file:test.db?cache=shared&mode=memory
	//:memory:
	//file::memory:

	// todo db path

	var err error
	DB, err = xorm.NewEngine("sqlite3", `file:C:\Users\shellus\go\src\github.com\shellus\my-telegram-bot\bin\telegram.sqlite3`)
	if err != nil {
		panic(err)
	}
	err = DB.Sync2(new(User))
	if err != nil {
		panic(err)
	}
}
type User struct {
	Id int64 `xorm:"pk autoincr"`
	Chat_id int64 `xorm:"unique"`
	Chat_meta *tgbotapi.Chat `xorm:"json"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	Status string
}