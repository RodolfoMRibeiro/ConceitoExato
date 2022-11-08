package routes

import (
	"conceitoExato/adapter/middleware"
	"conceitoExato/core/controller"

	"github.com/gin-gonic/gin"
)

func ModuleRoutes(modules *gin.RouterGroup) {
	modules.GET("/:moduleName", middleware.IsAuthorized(controller.FindModuleByName))
}
