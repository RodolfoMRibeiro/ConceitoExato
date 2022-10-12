package repository

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/model"
	"conceitoExato/common/library"
	"conceitoExato/common/util"
)

type IGoalRepository interface {
	GetGoalById(goalId string) error
	GetAllGoals() ([]model.Goal, error)
}

type goalRepository struct {
	user  *model.Goal
	users *[]model.Goal
}

func NewGoalRepository() IGoalRepository {
	return &goalRepository{
		user:  &model.Goal{},
		users: &[]model.Goal{},
	}
}

func (repository goalRepository) GetGoalById(goalId string) error {
	couldNotFindGoalById := db.GetGormDB().
		Table(library.TB_GOAL).
		Where("id = ?", goalId).
		First(repository.user).
		Error

	if util.ContainsError(couldNotFindGoalById) {
		return couldNotFindGoalById
	}

	return nil
}

func (repository goalRepository) GetAllGoals() ([]model.Goal, error) {
	couldNotFindGoals := db.GetGormDB().
		Table(library.TB_GOAL).
		Find(repository.users).
		Error

	if util.ContainsError(couldNotFindGoals) {
		return []model.Goal{}, couldNotFindGoals
	}

	return *repository.users, nil
}
