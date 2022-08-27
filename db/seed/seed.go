package seed

import (
	"conceitoExato/db"
	"conceitoExato/db/model"
	"conceitoExato/library"

	"gorm.io/gorm"
)

func SeedDatabase() {
	populateWithGoalData(db.GetGormDB())
	populateWithUserData(db.GetGormDB())
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

func populateWithUserData(db *gorm.DB) {

	users := []model.User{
		{
			GoalId:   1,
			Email:    "rodolfomarqribeiro@gmail.com",
			Password: "12345fjdskl",
			Login:    "rodolfo",
			Fullname: "93102830912",
			Courses:  []*model.Course{},
			Model:    gorm.Model{},
		},
	}

	for _, user := range users {
		if err := db.Table(library.TB_USER).Create(&user).Error; err != nil {
			break
		}
	}
}
