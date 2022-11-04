package router

import (
	"conceitoExato/core/routes"

	"github.com/gin-gonic/gin"
)

func Avaible(r *gin.Engine) {

	webSite := r.Group("/api")
	{
		usersGroup := webSite.Group("/user")
		routes.UserRoutes(usersGroup)

		goalsGroup := webSite.Group("/goal")
		routes.GoalRoutes(goalsGroup)
	}
}
