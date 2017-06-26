package user

import (
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
)

func TestCreate(t *testing.T) {
	DB.DropTables(&User{})
	err := DB.CreateTables(&User{})
	if err != nil {
		panic(err)
	}
	user := &User{
		Id:10000,
		Chat_id:100000,
		Chat_meta:&tgbotapi.Chat{ID:100000, Type:"ok", Title:"ok", UserName:"", FirstName:"", LastName:"", AllMembersAreAdmins:false},
	}
	_, err = DB.Insert(user)
	if err != nil {
		panic(err)
	}
}

func TestRead(t *testing.T){

	user := &User{}

	has, err := DB.Where("id = ?", 10000).Get(user)
	if err != nil {
		t.Error(err)
	}

	t.Log(has)
	fmt.Printf("%#v", user.Chat_meta)

	//DB.DropTables(user)
}