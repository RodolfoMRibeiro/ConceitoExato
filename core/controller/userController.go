package controller

import (
	"conceitoExato/common/util"
	"conceitoExato/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	user, couldNotFindUser := service.FindUser(ctx)

	if util.ContainsError(couldNotFindUser) {
		util.HttpNotFoundMessage(ctx, couldNotFindUser)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": user})
}

func Create(ctx *gin.Context) {
	couldNotCreateUserError := service.CreateUser(ctx)

	if couldNotCreateUserError != nil {
		util.HttpBadRequestMessage(ctx, couldNotCreateUserError)
		return
	}

	ctx.JSON(http.StatusCreated, "User created sucessfully")
}

func Delete(ctx *gin.Context) {
	_, couldNotDeleteUserError := service.DeleteUser(ctx)

	if util.ContainsError(couldNotDeleteUserError) {
		util.HttpNotFoundMessage(ctx, couldNotDeleteUserError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": "user was deleted successfuly"})
}

func ValidateLogin(ctx *gin.Context) {
	validatedLogin, couldNotValidateUserLogin := service.ValidateUserLogin(ctx)

	if util.ContainsError(couldNotValidateUserLogin) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": couldNotValidateUserLogin,
			"accepted_login": validatedLogin})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": nil, "accepted_login": validatedLogin})
}
