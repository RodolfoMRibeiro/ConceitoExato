package service

import (
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/util"
	"conceitoExato/core/goal/repository"

	"github.com/gin-gonic/gin"
)

func GetAllGoals(ctx *gin.Context) ([]model.Goal, error) {
	goalRepository := repository.NewGoalRepository()
	goals, couldNotGetAllGoals := goalRepository.GetAllGoals()

	if util.ContainsError(couldNotGetAllGoals) {
		return []model.Goal{}, nil
	}

	return goals, nil
}
