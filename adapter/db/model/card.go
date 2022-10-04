package model

import (
	"conceitoExato/common/library"

	"gorm.io/gorm"
)

type Card struct {
	ModuleId uint   `gorm:"column:module_id"`
	Name     string `gorm:"column:name"`
	Content  string `gorm:"column:content"`
	Kind     string `gorm:"column:kind"`
	gorm.Model
}

func (Card) TableName() string {
	return library.TB_CARD
}
