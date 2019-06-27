package users

import (
	"golang-cell5/src/pkg/db"
)

type Persons []Person

type Person db.Persons

func (u *Person) TableName() string {
	return "persons"
}
