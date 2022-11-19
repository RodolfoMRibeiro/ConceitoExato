package controller

import (
	"conceitoExato/common/util"
	"conceitoExato/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGoals(ctx *gin.Context) {
	responseGoals, couldNotGetAllGoals := service.GetAllGoals(ctx)

	if util.ContainsError(couldNotGetAllGoals) {
		util.HttpNotFoundMessage(ctx, couldNotGetAllGoals)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Goals": responseGoals})
}

func GetGoalByName(ctx *gin.Context) {
	responseGoal, couldNotGetGoalByName := service.GetGoalByName(ctx)

	if util.ContainsError(couldNotGetGoalByName) {
		util.HttpNotFoundMessage(ctx, couldNotGetGoalByName)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Goal": responseGoal})
}
