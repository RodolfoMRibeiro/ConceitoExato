package service

import (
	"conceitoExato/common/util"
	dto "conceitoExato/core/domain"
	"conceitoExato/core/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func FindByName(ctx *gin.Context) ([]dto.ModuleDto, error) {
	goalRepository := repository.NewModuleRepository()
	responseGoals, couldNotFindModuleByName := goalRepository.FindByName(ctx.Param("moduleName"))

	if util.ContainsError(couldNotFindModuleByName) {
		return []dto.ModuleDto{}, couldNotFindModuleByName
	}

	response := []dto.ModuleDto{}
	copier.Copy(&response, responseGoals)

	return response, nil
}
