package main

import (
	dbp "dev-blog/db"
	"dev-blog/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setUpDB() {
	db := dbp.New()
	dbp.SetUpDB(db)
}

func main() {
	setUpDB()
	routes.SetupRoutes()
}
