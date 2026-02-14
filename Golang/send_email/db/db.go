package db

import (
	"database/sql"
	add_news_update_db "login/db/add_newupdate_db"
	"login/db/db_init"
	"login/db/indexes"
	"login/db/newsletter_db"
	subscription_db "login/db/subscription_db"
)

func Initialize() {
	// Initialize database connection
	db_init.InitDB()

	database := db_init.GetDB()

	// Initialize all tables
	add_news_update_db.InitTable(database)
	newsletter_db.InitTable(database)
	subscription_db.InitTable(database)

	// Register all indexes
	indexes.RegisterIndexes(database)
}

func GetDB() *sql.DB {
	return db_init.GetDB()
}
