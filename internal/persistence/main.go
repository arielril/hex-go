package persistence

import (
	"sync"

	"hexgo/internal/types"
	. "hexgo/internal/types"

	"github.com/jinzhu/gorm"
)

type persistence struct {
	*gorm.DB
}

var db *gorm.DB
var once sync.Once

func getConn() string {
	return "root:root@(localhost:3306)/hexgo?charset=utf8"
}

func New() (types.Persistence, error) {
	var err error

	once.Do(func() {
		db, err = gorm.Open("mysql", getConn())
	})

	return persistence{db}, err
}

func (p persistence) Create(model interface{}) Persistence {
	db.Create(model)
	return p
}

func (p persistence) Find(model interface{}, where ...interface{}) Persistence {
	return p
}

func (p persistence) Update(model interface{}, where ...interface{}) Persistence {
	return p
}

func (p persistence) Delete(model interface{}, where ...interface{}) Persistence {
	return p
}
