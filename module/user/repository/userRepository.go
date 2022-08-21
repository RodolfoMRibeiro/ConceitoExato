package repository

import (
	"conceitoExato/db"
	"conceitoExato/db/model"
	"conceitoExato/util"
)

type IUserRepository interface {
	FindUserByLogin(string) error
	FindUser() error
}

type userRepository struct {
	user  model.User
	users []model.User
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		user:  model.User{},
		users: []model.User{},
	}
}

func (repository *userRepository) FindUserByLogin(login string) error {
	couldNotFindUserLogin := db.GetGormDB().Where("login = ?", login).First(repository.user).Error

	if util.ContainsError(couldNotFindUserLogin) {
		return couldNotFindUserLogin
	}

	return nil
}

func (repository *userRepository) FindUser() error {
	couldNotFindUser := db.GetGormDB().First(repository.user).Error

	if util.ContainsError(couldNotFindUser) {
		return couldNotFindUser
	}

	return nil
}

func (repository userRepository) GetUser() model.User {
	return repository.user
}
