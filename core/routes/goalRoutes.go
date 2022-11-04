package routes

import (
	"conceitoExato/core/controller"

	"github.com/gin-gonic/gin"
)

func GoalRoutes(goal *gin.RouterGroup) {
	goal.GET("/", controller.GetAllGoals)
	goal.GET("/:name", controller.GetGoalByName)
}
