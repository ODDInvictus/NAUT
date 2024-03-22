package models

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var DB *sqlx.DB

func InitDB() *sqlx.DB {
	connString := viper.GetString("database.connection_string")

	if connString == "" {
		panic("database connection string (database.connection_string) not found")
	}

	log.Printf("Starting database connection to: %s\n", connString)

	db, err := sqlx.Open("mysql", connString)

	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
	return db
}