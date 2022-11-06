package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseHolder struct {
	db            *gorm.DB
	configuration string
}

func NewDatabase(config string) *databaseHolder {
	mysql := &databaseHolder{
		db:            &gorm.DB{},
		configuration: config,
	}
	return mysql
}

func (m *databaseHolder) connect() error {
	var err error
	m.db, err = gorm.Open(postgres.Open(m.configuration), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	return nil
}
