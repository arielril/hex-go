package mysql

import (
	"sync"

	"github.com/jinzhu/gorm"
	// Get the mysql dialect to connect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Persistence hold the persistence
type Persistence struct {
	*gorm.DB
}

// Model of the database
type Model = gorm.Model

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

// NewConnection create a new connection to the database
func NewConnection() (*Persistence, error) {
	database, err := initDatabase()

	return &Persistence{database}, err
}
