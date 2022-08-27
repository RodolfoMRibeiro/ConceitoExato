package controller

import (
	"conceitoExato/module/user/repository"
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
	repo := repository.NewUserRepository()

	bodyRequest, unableToGetRawData := ctx.GetRawData()

	if util.ContainsError(unableToGetRawData) {
		util.HttpBadRequestMessage(ctx, unableToGetRawData)
		return
	}

	couldNotCreateUser := repo.CreateUser(bodyRequest)

	if util.ContainsError(couldNotCreateUser) {
		util.HttpBadRequestMessage(ctx, couldNotCreateUser)
		return
	}
}

func Delete(ctx *gin.Context) {
	repo := repository.NewUserRepository()

	couldNotDeleteUser := repo.DeleteUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotDeleteUser) {
		util.HttpNotFoundMessage(ctx, couldNotDeleteUser)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": repo.GetUser()})
}
