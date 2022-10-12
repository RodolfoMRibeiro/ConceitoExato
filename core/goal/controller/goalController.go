package controller

import (
	"conceitoExato/common/util"
	"conceitoExato/core/goal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGoals(ctx *gin.Context) {
	goals, couldNotGetAllGoals := service.GetAllGoals(ctx)

	if util.ContainsError(couldNotGetAllGoals) {
		util.HttpNotFoundMessage(ctx, couldNotGetAllGoals)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Goals": goals})
}
