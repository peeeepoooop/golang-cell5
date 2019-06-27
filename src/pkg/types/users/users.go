package users

import (
	"golang-cell5/src/pkg/db"
)

type Users []User

type User db.Users

func (u *User) TableName() string {
	return "users"
}
