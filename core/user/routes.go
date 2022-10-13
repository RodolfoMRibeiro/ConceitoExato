package user

import (
	"conceitoExato/adapter/middleware"
	"conceitoExato/core/user/controller"

	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.POST("/token/:username", middleware.GenerateJWT)
	users.GET("/:login", middleware.IsAuthorized(controller.Find))
	users.POST("/", controller.Create)
	users.POST("/login", controller.ValidateLogin)
	users.DELETE("/:login", controller.Delete)
	// users.PATCH("/", controller.Update)
}
