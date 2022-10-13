package model

import (
	"conceitoExato/common/library"

	"gorm.io/gorm"
)

type Course struct {
	Name          string   `gorm:"column:name"             json:"name"`
	Description   string   `gorm:"column:description"      json:"description"`
	ModulesAmount uint     `gorm:"column:modules_amount"   json:"modules_amount"`
	Modules       []Module `gorm:"foreignKey:course_id"    json:"modules"`
	Users         []*User  `gorm:"many2many:user_course"   json:"users"`
	gorm.Model
}

func (Course) TableName() string {
	return library.TB_COURSE
}
