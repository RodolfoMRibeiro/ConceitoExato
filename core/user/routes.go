package user

import (
	"conceitoExato/core/user/controller"

	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/:login", controller.Find)
	users.POST("/", controller.Create)
	users.POST("/login", controller.ValidateLogin)
	users.DELETE("/:login", controller.Delete)
	// users.PATCH("/", controller.Update)
}
