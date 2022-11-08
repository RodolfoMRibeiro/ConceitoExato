package interfaces

import "conceitoExato/adapter/db/model"

type IModuleInterface interface {
	FindByName(string) (*model.Module, error)
}
