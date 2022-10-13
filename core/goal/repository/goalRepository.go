package repository

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/library"
	"conceitoExato/common/util"
)

type IGoalRepository interface {
	GetGoalByName(string) (model.Goal, error)
	GetAllGoals() ([]model.Goal, error)
}

type goalRepository struct {
	goal  *model.Goal
	goals *[]model.Goal
}

func NewGoalRepository() IGoalRepository {
	return &goalRepository{
		goal:  &model.Goal{},
		goals: &[]model.Goal{},
	}
}

func (repository goalRepository) GetGoalByName(goalName string) (model.Goal, error) {
	couldNotFindGoalById := db.GetGormDB().
		Table(library.TB_GOAL).
		Where("name = ?", goalName).
		First(repository.goal).
		Error

	if util.ContainsError(couldNotFindGoalById) {
		return model.Goal{}, couldNotFindGoalById
	}

	return *repository.goal, nil
}

func (repository goalRepository) GetAllGoals() ([]model.Goal, error) {
	couldNotFindGoals := db.GetGormDB().
		Table(library.TB_GOAL).
		Find(repository.goals).
		Error

	if util.ContainsError(couldNotFindGoals) {
		return []model.Goal{}, couldNotFindGoals
	}

	return *repository.goals, nil
}
