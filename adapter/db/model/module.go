package model

import (
	"conceitoExato/common/library"

	"gorm.io/gorm"
)

type Module struct {
	CourseId    uint   `gorm:"column:course_id"         json:"course_id"`
	CardsAmount uint   `gorm:"column:cards_amount"      json:"cards_amount"`
	Name        string `gorm:"column:name"              json:"name"`
	Description string `gorm:"column:description"       json:"description"`
	Cards       []Card `gorm:"foreignKey:module_id"     json:"cards"`
	gorm.Model
}

func (Module) TableName() string {
	return library.TB_MODULE
}
