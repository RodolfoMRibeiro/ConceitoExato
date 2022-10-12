package goal

import (
	"conceitoExato/core/goal/controller"

	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(goal *gin.RouterGroup) {
	goal.GET("/goal", controller.GetAllGoals)
}
