package interfaces

import "conceitoExato/adapter/db/model"

type IUserRepository interface {
	FindUserByLogin(string) error
	CreateUser([]byte) error
	DeleteUserByLogin(string) error
	IsValidateUser(string, string) bool
	GetUser() *model.User
	SetUser([]byte) error
}
