package user

import (
	"testing"
	_ "github.com/mattn/go-sqlite3"
)

func TestCreate(t *testing.T) {
	//js := `{"id":249603346,"type":"private","title":"","username":"","first_name":"shellus","last_name":"","all_members_are_administrators":false}`
	//tgChat := &tgbotapi.Chat{}
	//err := json.Unmarshal([]byte(js), tgChat)
	//if err != nil {
	//	panic(err)
	//}
}

func TestRead(t *testing.T){
	user := &User{Id:3}
	err := DB.Read(user)
	if err != nil{
		t.Error(err)
	}
	t.Log(user.Chat())
}