package domain

import "conceitoExato/adapter/db/model"

type Course struct {
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	ModulesAmount uint           `json:"modules_amount"`
	Modules       []model.Module `json:"modules"`
}
