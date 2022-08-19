package model

type Goal struct {
	Id   string `gorm:"id; primaryKey"`
	Name string
}

type User struct {
	Id       int  `gorm:"id; primaryKey"`
	Goal     Goal `gorm:"foreignKey:Id"`
	Email    string
	Password string
	Login    string
	Fullname string
}

type Course struct {
	Id            int `gorm:"id; primaryKey"`
	Name          string
	Description   string
	ModulesAmount int
}

type CourseInProgress struct {
	Courses []Course
	Users   []User
}

type Module struct {
	Id          int    `gorm:"id; primaryKey"`
	CourseId    Course `gorm:"foreignKey:Id"`
	CardsAmount int
	Name        string
	Description string
}

type Card struct {
	Id       int    `gorm:"id; primaryKey"`
	ModuleId Module `gorm:"foreignKey:Id"`
	Name     string
	Content  string
	Kind     string
}
