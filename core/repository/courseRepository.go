package repository

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/library"
	"conceitoExato/common/util"
	"conceitoExato/core/interfaces"
)

type courseRepository struct {
	course  model.Course
	courses []model.Course
}

func NewCourseRepository() interfaces.ICourseRepository {
	return &courseRepository{
		course:  model.Course{},
		courses: []model.Course{},
	}
}

func (repository *courseRepository) GetAllCourses() ([]model.Course, error) {
	couldNotFindAnyCourse := db.GetGormDB().
		Table(library.TB_COURSE).
		Find(&repository.courses).
		Error

	if util.ContainsError(couldNotFindAnyCourse) {
		return repository.courses, couldNotFindAnyCourse
	}

	return repository.courses, nil
}
