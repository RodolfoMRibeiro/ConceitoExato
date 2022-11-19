package controller

import (
	"conceitoExato/common/util"
	"conceitoExato/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCourses(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	courses, couldNotGetAllCourses := service.GetAllCourses(ctx)

	if util.ContainsError(couldNotGetAllCourses) {
		util.HttpNotFoundMessage(ctx, couldNotGetAllCourses)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Courses": courses})
}

func CreateCourse(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	couldNotCreateCourse := service.CreateCourse(ctx)

	if util.ContainsError(couldNotCreateCourse) {
		util.HttpBadRequestMessage(ctx, couldNotCreateCourse)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": couldNotCreateCourse})
}
