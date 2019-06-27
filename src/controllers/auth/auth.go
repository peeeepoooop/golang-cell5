package session

import (
	"github.com/go-xorm/xorm"
	Users "golang-cell5/src/pkg/types/users"
)

var db *xorm.Engine

type LoginData struct {
	Token string     `json:"token"`
	User  Users.User `json:"user"`
}

func Init(DB *xorm.Engine) {
	db = DB
}
