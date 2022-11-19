package controller

import (
	"conceitoExato/common/util"
	"conceitoExato/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindModuleByName(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	module, couldNotGetModuleByName := service.FindByName(ctx)

	if util.ContainsError(couldNotGetModuleByName) {
		util.HttpNotFoundMessage(ctx, couldNotGetModuleByName)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Module": module})
}
