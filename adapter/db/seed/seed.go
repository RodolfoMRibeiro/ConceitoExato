package seed

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/library"

	"gorm.io/gorm"
)

func SeedDatabase() {
	populateWithGoalData(db.GetGormDB())
}

func populateWithGoalData(db *gorm.DB) {

	goals := []model.Goal{
		{Model: gorm.Model{}, Users: []model.User{}, Name: "Teste"},
		{Model: gorm.Model{}, Users: []model.User{}, Name: "TesteAgain"},
	}

	for _, goal := range goals {
		if err := db.Table(library.TB_GOAL).Create(&goal).Error; err != nil {
			break
		}
	}
}
