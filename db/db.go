package db

import (
	"dev-blog/models"
	"dev-blog/utils"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // register postgres driver
)

// These would actually come from enviroment variables.
// An example of that is available here:
const (
	user   = "tcfbvjkz"
	pwd    = "wW9I828Tn8S206Fkkd9q-cJJfpOcYHWD"
	url    = "isilo.db.elephantsql.com"
	port   = "5432"
	dbName = "tcfbvjkz"
)

// SetUpDB Resets database and migrate table users schema
func SetUpDB(db *gorm.DB) {
	fmt.Println("Resetting database")
	db.DropTableIfExists(&models.User{})
	db.AutoMigrate(&models.User{})
}

// New creates a database connection
func New() *gorm.DB {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pwd, url, port, dbName)
	db, err := gorm.Open("postgres", dbConnString)
	utils.Must(err)
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	err = db.DB().Ping()
	utils.Must(err)
	return db
}
