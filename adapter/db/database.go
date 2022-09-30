package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDb struct {
	db            *gorm.DB
	configuration string
}

func NewMysql(config string) *mysqlDb {
	mysql := &mysqlDb{
		db:            &gorm.DB{},
		configuration: config,
	}
	return mysql
}

func (m *mysqlDb) connect() error {
	var err error
	m.db, err = gorm.Open(mysql.Open(m.configuration), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	return nil
}
