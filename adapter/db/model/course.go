package model

import (
	"conceitoExato/common/library"

	"gorm.io/gorm"
)

type Course struct {
	Name          string   `gorm:"column:name"`
	Description   string   `gorm:"column:description"`
	ModulesAmount uint     `gorm:"column:modules_amount"`
	Modules       []Module `gorm:"foreignKey:course_id"`
	Users         []*User  `gorm:"many2many:user_course;"`
	gorm.Model
}

func (Course) TableName() string {
	return library.TB_COURSE
}
