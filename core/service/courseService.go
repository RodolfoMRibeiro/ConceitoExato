package service

import (
	"conceitoExato/common/util"
	dto "conceitoExato/core/domain"
	"conceitoExato/core/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetAllCourses(ctx *gin.Context) (dto.Course, error) {
	courseRepository := repository.NewCourseRepository()

	courses, couldNotGetCourses := courseRepository.GetAllCourses()

	if util.ContainsError(couldNotGetCourses) {
		return dto.Course{}, couldNotGetCourses
	}

	courseDto := dto.Course{}
	copier.Copy(&courseDto, courses)

	return courseDto, nil
}
