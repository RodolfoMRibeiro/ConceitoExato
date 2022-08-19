package model

import "conceitoExato/library"

type Goal struct {
	Id   string `gorm:"id; primaryKey"`
	Name string
}

func (Goal) TableName() string {
	return library.TB_GOAL
}

type User struct {
	Id       int  `gorm:"id; primaryKey"`
	Goal     Goal `gorm:"foreignKey:Id"`
	Email    string
	Password string
	Login    string
	Fullname string
}

func (User) TableName() string {
	return library.TB_USER
}

type Course struct {
	Id            int `gorm:"id; primaryKey"`
	Name          string
	Description   string
	ModulesAmount int
}

func (Course) TableName() string {
	return library.TB_COURSE
}

type CourseInProgress struct {
	Courses []Course
	Users   []User
}

func (CourseInProgress) TableName() string {
	return library.TB_COURSE_IN_PROGRESS
}

type Module struct {
	Id          int    `gorm:"id; primaryKey"`
	CourseId    Course `gorm:"foreignKey:Id"`
	CardsAmount int
	Name        string
	Description string
}

func (Module) TableName() string {
	return library.TB_MODULE
}

type Card struct {
	Id       int    `gorm:"id; primaryKey"`
	ModuleId Module `gorm:"foreignKey:Id"`
	Name     string
	Content  string
	Kind     string
}

func (Card) TableName() string {
	return library.TB_CARD
}
