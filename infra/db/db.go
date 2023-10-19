package db

import (
	"log"

	"github.com/emiliosheinz/imersao-full-cycle-challenge-01/domain/model"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	log.Println("Connecting to database...")
}

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&model.Product{})

	return db
}
