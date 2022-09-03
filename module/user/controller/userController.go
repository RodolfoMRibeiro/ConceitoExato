package controller

import (
	"conceitoExato/module/user/repository"
	"conceitoExato/module/user/service"
	"conceitoExato/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	repo := repository.NewUserRepository()

	couldNotFindUser := repo.FindUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotFindUser) {
		util.HttpNotFoundMessage(ctx, couldNotFindUser)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": repo.GetUser()})
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

	createdUser, couldNotDeleteUserError := service.DeleteUser(ctx)

	if util.ContainsError(couldNotDeleteUserError) {
		util.HttpNotFoundMessage(ctx, couldNotDeleteUserError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": createdUser})
}
