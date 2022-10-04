package model

import (
	"conceitoExato/common/library"

	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	Users []User `gorm:"foreignKey:goal_id"`
	Name  string `gorm:"column:name"`
}

func (Goal) TableName() string {
	return library.TB_GOAL
}
