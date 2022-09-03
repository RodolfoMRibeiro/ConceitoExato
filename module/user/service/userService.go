package service

import (
	"conceitoExato/db/model"
	"conceitoExato/module/user/repository"
	"conceitoExato/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) error {
	repo := repository.NewUserRepository()

	bodyRequest, unableToGetRawData := ctx.GetRawData()

	if util.ContainsError(unableToGetRawData) {
		return unableToGetRawData
	}

	couldNotCreateUser := repo.CreateUser(bodyRequest)

	if util.ContainsError(couldNotCreateUser) {
		return couldNotCreateUser
	}

	return nil
}

func FindUser(ctx *gin.Context) (*model.User, error) {
	repo := repository.NewUserRepository()

	couldNotFindUser := repo.FindUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotFindUser) {
		return &model.User{}, nil
	}

	return repo.GetUser(), nil
}

func DeleteUser(ctx *gin.Context) (*model.User, error) {
	repo := repository.NewUserRepository()

	couldNotDeleteUser := repo.DeleteUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotDeleteUser) {
		return &model.User{}, couldNotDeleteUser
	}

	return repo.GetUser(), nil
}
