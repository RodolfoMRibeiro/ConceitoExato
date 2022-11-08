package repository

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/library"
	"conceitoExato/common/util"
	"conceitoExato/core/interfaces"
)

type moduleRepository struct {
	module  *model.Module
	modules *[]model.Module
}

func NewModuleRepository() interfaces.IModuleInterface {
	return &moduleRepository{
		module:  &model.Module{},
		modules: &[]model.Module{},
	}
}

func (repository *moduleRepository) FindByName(moduleName string) (*model.Module, error) {
	couldNotFindModuleByName := db.GetGormDB().
		Table(library.TB_MODULE).
		Where("name = ?", moduleName).
		Find(repository.module).
		Error

	if util.ContainsError(couldNotFindModuleByName) {
		return repository.module, couldNotFindModuleByName
	}

	return repository.module, nil
}
