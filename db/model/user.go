package model

import (
	"conceitoExato/library"

	"gorm.io/gorm"
)

type IUser interface {
}

type User struct {
	gorm.Model
	GoalId   uint      `gorm:"column:goal_id"`
	Email    string    `gorm:"column:email"`
	Password string    `gorm:"column:password"`
	Login    string    `gorm:"column:login"`
	Fullname string    `gorm:"column:fullname"`
	Courses  []*Course `gorm:"many2many:user_course;"`
}

func (User) TableName() string {
	return library.TB_USER
}
