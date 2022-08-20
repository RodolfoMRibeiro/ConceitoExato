package db

import (
	"conceitoExato/env"
	"fmt"
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func GetGormDB() *gorm.DB {
	return db
}

func StartDatabase() {
	databaseConfiguration := createDatabaseStringConfig()
	mysql := NewMysql(databaseConfiguration)

	if err := mysql.connect(); err != nil {
		log.Fatal(err)
	}

	loadMigrations(mysql.db)

	db = mysql.db
}

func createDatabaseStringConfig() string {
	databaseStringConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s%s",

		env.Mysql.USER,
		env.Mysql.PASSWORD,
		env.Mysql.HOST,
		env.Mysql.PORT,
		env.Mysql.DB,
		env.Mysql.ADDITIONAL_CONFIGS,
	)

	return databaseStringConfig
}

func loadMigrations(db *gorm.DB) {
}
