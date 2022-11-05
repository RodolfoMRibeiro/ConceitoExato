package interfaces

import "conceitoExato/adapter/db/model"

type ICourseRepository interface {
	GetAllCourses() ([]model.Course, error)
}
