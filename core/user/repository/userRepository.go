package repository

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/util"
	"encoding/json"
)

type IUserRepository interface {
	FindUserByLogin(string) error
	CreateUser([]byte) error
	DeleteUserByLogin(string) error
	GetUser() *model.User
	SetUser([]byte) error
}

type userRepository struct {
	user  *model.User
	users *[]model.User
}

func NewUserRepository() IUserRepository {
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

	if util.ContainsError(couldNotUnmarshal) {
		return couldNotUnmarshal
	}

	couldNotCreateUser := db.GetGormDB().Create(&repository.user).Error

	if util.ContainsError(couldNotCreateUser) {
		return couldNotCreateUser
	}

	return nil
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
