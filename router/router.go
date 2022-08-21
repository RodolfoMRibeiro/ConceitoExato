package router

import (
	"conceitoExato/module/user"

	"github.com/gin-gonic/gin"
)

func Avaible(r *gin.Engine) {

	webSite := r.Group("/api")
	{
		usersGroup := webSite.Group("/user")
		user.AvaiableRoutes(usersGroup)
	}
}
