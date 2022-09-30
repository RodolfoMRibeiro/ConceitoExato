package service

import (
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/util"
	"conceitoExato/core/user/repository"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) error {
	userRepository := repository.NewUserRepository()

	bodyRequest, unableToGetRawData := ctx.GetRawData()

	if util.ContainsError(unableToGetRawData) {
		return unableToGetRawData
	}

	couldNotCreateUser := userRepository.CreateUser(bodyRequest)

	if util.ContainsError(couldNotCreateUser) {
		return couldNotCreateUser
	}

	return nil
}

func FindUser(ctx *gin.Context) (*model.User, error) {
	userRepository := repository.NewUserRepository()

	couldNotFindUser := userRepository.FindUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotFindUser) {
		return &model.User{}, nil
	}

	return userRepository.GetUser(), nil
}

func DeleteUser(ctx *gin.Context) (*model.User, error) {
	userRepository := repository.NewUserRepository()

	couldNotDeleteUser := userRepository.DeleteUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotDeleteUser) {
		return &model.User{}, couldNotDeleteUser
	}

	return userRepository.GetUser(), nil
}
