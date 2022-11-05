package repository

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/library"
	"conceitoExato/common/util"
	"conceitoExato/core/interfaces"
	"encoding/json"
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

func (repository courseRepository) CreateCourse(bodyInBytes []byte) error {
	json.Unmarshal(bodyInBytes, &repository.course)

	couldNotCreateUser := db.GetGormDB().
		Table(library.TB_COURSE).
		Create(&repository.course).
		Error

	if util.ContainsError(couldNotCreateUser) {
		return couldNotCreateUser
	}

	return nil
}
