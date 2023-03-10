package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func NewDB() (*sql.DB, error) {

	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// "root:Kifeb99##@tcp(localhost:3306)/ecommerce_api"
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}
