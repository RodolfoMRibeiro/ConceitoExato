package routes

import (
	"conceitoExato/adapter/middleware"
	"conceitoExato/core/controller"

	"github.com/gin-gonic/gin"
)

func CourseRoutes(course *gin.RouterGroup) {
	course.GET("/allCourses", middleware.IsAuthorized(controller.GetAllCourses))
}
