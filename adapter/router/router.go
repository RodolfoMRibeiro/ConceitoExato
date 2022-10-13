package router

import (
	"conceitoExato/core/goal"
	"conceitoExato/core/user"

	"github.com/gin-gonic/gin"
)

func Avaible(r *gin.Engine) {

	webSite := r.Group("/api")
	{
		usersGroup := webSite.Group("/user")
		user.AvaiableRoutes(usersGroup)

		goalsGroup := webSite.Group("/goal")
		goal.AvaiableRoutes(goalsGroup)
	}
}
