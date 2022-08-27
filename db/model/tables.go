package model

import (
	"conceitoExato/library"

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
