package routes

import (
	"conceitoExato/adapter/middleware"
	"conceitoExato/core/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(users *gin.RouterGroup) {
	users.POST("/token/generate", middleware.GenerateJWT)
	users.GET("/:login", middleware.IsAuthorized(controller.Find))
	users.POST("/create", controller.Create)
	users.POST("/login/validate", controller.ValidateLogin)
	users.DELETE("/:login", middleware.IsAuthorized(controller.Delete))
}
