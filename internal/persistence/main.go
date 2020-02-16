package persistence

import (
	"sync"

	. "github.com/arielril/hexgo/internal/types"

	"github.com/jinzhu/gorm"
)

type Persistence interface {
	Create(model interface{}) Persistence
	Find(model interface{}, where ...interface{}) FindResult
	Update(model interface{}, where ...interface{}) Persistence
	Delete(model interface{}, where ...interface{}) Persistence
}

type persistence struct {
	*gorm.DB
}

var db *gorm.DB
var once sync.Once

func getConn() string {
	return "root:root@(localhost:3306)/hexgo?charset=utf8"
}

func initDatabase() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		db, err = gorm.Open("mysql", getConn())
	})

	return db, err
}

func New() (Persistence, error) {
	database, err := initDatabase()

	return &persistence{database}, err
}

func (p *persistence) Create(model interface{}) Persistence {
	db.Create(model)
	return p
}

func (p *persistence) Find(model interface{}, where ...interface{}) FindResult {
	return FindResult{
		P:     p,
		Value: p.DB.Find(model, where).Value,
	}
}

func (p *persistence) Update(model interface{}, where ...interface{}) Persistence {
	return p
}

func (p *persistence) Delete(model interface{}, where ...interface{}) Persistence {
	return p
}
