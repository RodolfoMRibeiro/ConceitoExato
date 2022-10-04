package model

import (
	"conceitoExato/common/library"

	"gorm.io/gorm"
)

type Module struct {
	CourseId    uint   `gorm:"column:course_id"`
	CardsAmount uint   `gorm:"column:cards_amount"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Cards       []Card `gorm:"foreignKey:module_id"`
	gorm.Model
}

func (Module) TableName() string {
	return library.TB_MODULE
}
