package persons

import (
	"github.com/go-xorm/xorm"
)

func FindBy(db *xorm.Engine, limit int, offset int) (persons Persons, err error) {
	err = db.
		Limit(limit, offset).
		Find(&persons)

	return
}
