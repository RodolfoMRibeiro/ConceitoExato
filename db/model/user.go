package model

import (
	"conceitoExato/library"

	"gorm.io/gorm"
)

type IUser interface {
}

type User struct {
	GoalId   uint      `gorm:"column:goal_id"           json:"goal_id"`
	Email    string    `gorm:"column:email"             json:"email"`
	Password string    `gorm:"column:password"          json:"password"`
	Login    string    `gorm:"column:login; unique"     json:"login"`
	Fullname string    `gorm:"column:fullname"          json:"fullname"`
	Courses  []*Course `gorm:"many2many:user_course;"   json:"courses"`
	gorm.Model
}

func (User) TableName() string {
	return library.TB_USER
}
