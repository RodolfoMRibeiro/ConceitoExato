package db

import (
	"conceitoExato/adapter/db/model"
	"conceitoExato/adapter/env"
	"fmt"
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func GetGormDB() *gorm.DB {
	return db
}

func StartDatabase() {
	postgresDB := NewDatabase(env.Database.DATABASE_URL)

	if err := postgresDB.connect(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	loadMigrations(postgresDB.db)

	db = postgresDB.db

	fmt.Println("Connected to Database sucessfully")
}

func loadMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Goal{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Course{})
	db.AutoMigrate(&model.Module{})
	db.AutoMigrate(&model.Card{})
}
