package user

import (
	"conceitoExato/module/user/controller"

	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/:login", controller.Find)
	users.POST("/", controller.Create)
	users.DELETE("/:login", controller.Delete)
	// users.PATCH("/", controller.Update)
}
