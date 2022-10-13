package model

import (
	"conceitoExato/common/library"

	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	Users []User `gorm:"foreignKey:goal_id"   json:"users"`
	Name  string `gorm:"column:name; unique"  json:"name"`
}

func (Goal) TableName() string {
	return library.TB_GOAL
}
