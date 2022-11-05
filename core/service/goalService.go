package service

import (
	"conceitoExato/common/util"
	dto "conceitoExato/core/domain"
	"conceitoExato/core/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetAllGoals(ctx *gin.Context) ([]dto.GoalDto, error) {
	goalRepository := repository.NewGoalRepository()
	responseGoals, couldNotGetAllGoals := goalRepository.GetAllGoals()

	if util.ContainsError(couldNotGetAllGoals) {
		return []dto.GoalDto{}, couldNotGetAllGoals
	}

	response := []dto.GoalDto{}
	copier.Copy(&response, responseGoals)

	return response, nil
}

func GetGoalByName(ctx *gin.Context) (dto.GoalDto, error) {
	goalRepository := repository.NewGoalRepository()
	goalFromDatabase, couldNotGetGoalByName := goalRepository.GetGoalByName(ctx.Param("name"))

	if util.ContainsError(couldNotGetGoalByName) {
		return dto.GoalDto{}, couldNotGetGoalByName
	}

	response := dto.GoalDto{}
	copier.Copy(&response, goalFromDatabase)

	return response, nil
}
