package config

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Connect establishes a connection to the database.
func Connect() {
	username := "root"
	password := "ayushsql"
	host := "localhost"
	port := "3306"
	dbname := "bookstore"
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	// Open a connection to the database
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Log a message indicating successful database connection
	log.Println("Connected to the database")
}

// GetDB returns the established database connection.
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection not initialized. Call Connect() first.")
	}
	return db
}