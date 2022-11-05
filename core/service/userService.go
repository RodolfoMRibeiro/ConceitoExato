package service

import (
	"conceitoExato/adapter/db/model"
	"conceitoExato/adapter/middleware"
	"conceitoExato/common/util"
	dto "conceitoExato/core/domain"
	"conceitoExato/core/repository"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func CreateUser(ctx *gin.Context) error {
	userRepository := repository.NewUserRepository()
	requestUser := model.User{}

	if unableToBindJson := ctx.BindJSON(&requestUser); util.ContainsError(unableToBindJson) {
		return unableToBindJson
	}

	hashedPassword, hashPasswordError := middleware.HashPassword(requestUser.Password)

	if util.ContainsError(hashPasswordError) {
		return hashPasswordError
	}

	requestUser.Password = hashedPassword
	bodyRequestInBytes, couldNotUnmarshalStruct := json.Marshal(requestUser)

	if util.ContainsError(couldNotUnmarshalStruct) {
		return couldNotUnmarshalStruct
	}

	couldNotCreateUser := userRepository.CreateUser(bodyRequestInBytes)

	if util.ContainsError(couldNotCreateUser) {
		return couldNotCreateUser
	}

	return nil
}

func FindUser(ctx *gin.Context) (dto.UserDto, error) {
	userRepository := repository.NewUserRepository()
	couldNotFindUser := userRepository.FindUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotFindUser) {
		return dto.UserDto{}, nil
	}

	userFromRepository := userRepository.GetUser()
	userDto := dto.UserDto{}

	copier.Copy(&userDto, userFromRepository)

	return userDto, nil
}

func DeleteUser(ctx *gin.Context) (*model.User, error) {
	userRepository := repository.NewUserRepository()
	couldNotDeleteUser := userRepository.DeleteUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotDeleteUser) {
		return &model.User{}, couldNotDeleteUser
	}

	return userRepository.GetUser(), nil
}

func ValidateUserLogin(ctx *gin.Context) (bool, error) {
	userRepository := repository.NewUserRepository()
	requestUser := model.User{}

	if unableToBindJson := ctx.BindJSON(&requestUser); util.ContainsError(unableToBindJson) {
		return false, unableToBindJson
	}

	couldNotFindUser := userRepository.FindUserByLogin(requestUser.Login)

	if util.ContainsError(couldNotFindUser) || userRepository.GetUser().IsEmpty() {
		return false, errors.New("user could not be found")
	}

	if userRepository.IsValidateUser(requestUser.Login, requestUser.Password) {
		return true, nil
	}

	return false, errors.New("user is not the same")
}
