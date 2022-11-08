package router

import (
	"conceitoExato/core/routes"

	"github.com/gin-gonic/gin"
)

func Avaible(r *gin.Engine) {

	apiV1 := r.Group("/api")
	{
		usersGroup := apiV1.Group("/user")
		routes.UserRoutes(usersGroup)

		goalsGroup := apiV1.Group("/goal")
		routes.GoalRoutes(goalsGroup)

		coursesGroup := apiV1.Group("/course")
		routes.CourseRoutes(coursesGroup)

		modulesGroup := apiV1.Group("/module")
		routes.ModuleRoutes(modulesGroup)
	}
}
