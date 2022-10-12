package service

import (
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/util"
	"conceitoExato/core/goal/repository"

	"github.com/gin-gonic/gin"
)

func GetAllGoals(ctx *gin.Context) ([]model.Goal, error) {
	goalRepository := repository.NewGoalRepository()
	responseGoals, couldNotGetAllGoals := goalRepository.GetAllGoals()

	if util.ContainsError(couldNotGetAllGoals) {
		return []model.Goal{}, couldNotGetAllGoals
	}

	return responseGoals, nil
}

func GetGoalByName(ctx *gin.Context) (model.Goal, error) {
	goalRepository := repository.NewGoalRepository()
	responseGoal, couldNotGetGoalByName := goalRepository.GetGoalByName(ctx.Param("name"))

	if util.ContainsError(couldNotGetGoalByName) {
		return model.Goal{}, couldNotGetGoalByName
	}

	return responseGoal, nil
}
