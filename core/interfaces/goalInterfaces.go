package interfaces

import "conceitoExato/adapter/db/model"

type IGoalRepository interface {
	GetGoalByName(string) (model.Goal, error)
	GetAllGoals() ([]model.Goal, error)
}
