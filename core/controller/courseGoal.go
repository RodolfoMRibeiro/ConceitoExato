package controller

import (
	"conceitoExato/common/util"
	"conceitoExato/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCourses(ctx *gin.Context) {
	courses, couldNotGetAllCourses := service.GetAllCourses(ctx)

	if util.ContainsError(couldNotGetAllCourses) {
		util.HttpNotFoundMessage(ctx, couldNotGetAllCourses)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Courses": courses})
}
