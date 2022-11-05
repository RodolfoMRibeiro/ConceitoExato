package repository

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/model"
	"conceitoExato/adapter/middleware"
	"conceitoExato/common/util"
	"conceitoExato/core/interfaces"
	"encoding/json"
)

type userRepository struct {
	user  *model.User
	users *[]model.User
}

func NewUserRepository() interfaces.IUserRepository {
	return &userRepository{
		user:  &model.User{},
		users: &[]model.User{},
	}
}

func (repository *userRepository) FindUserByLogin(login string) error {
	couldNotFindUserLogin := db.GetGormDB().
		Where("login = ?", login).
		Find(repository.user).
		Error

	if util.ContainsError(couldNotFindUserLogin) {
		return couldNotFindUserLogin
	}

	return nil
}

func (repository *userRepository) DeleteUserByLogin(login string) error {
	couldNotDeleteUser := db.GetGormDB().
		Where("login = ?", login).
		Delete(repository.user).
		Error

	if util.ContainsError(couldNotDeleteUser) {
		return couldNotDeleteUser
	}

	return nil
}

func (repository *userRepository) CreateUser(jsonElement []byte) error {
	couldNotUnmarshal := repository.SetUser(jsonElement)
	repository.user.GoalId = 1

	if util.ContainsError(couldNotUnmarshal) {
		return couldNotUnmarshal
	}

	couldNotCreateUser := db.GetGormDB().Create(&repository.user).Error

	if util.ContainsError(couldNotCreateUser) {
		return couldNotCreateUser
	}

	return nil
}

func (repository *userRepository) IsValidateUser(login, hashedPassword string) bool {
	repository.FindUserByLogin(login)

	if util.IsEqual(login, repository.GetUser().Login) &&
		middleware.CheckPasswordHash(hashedPassword, repository.GetUser().Password) {
		return true
	}

	return false
}

func (repository userRepository) GetUser() *model.User {
	return repository.user
}

func (repository *userRepository) SetUser(jsonElement []byte) error {
	couldNotUnmarshal := json.Unmarshal(jsonElement, &repository.user)

	if util.ContainsError(couldNotUnmarshal) {
		return couldNotUnmarshal
	}

	return nil
}
