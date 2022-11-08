package domain

type ModuleDto struct {
	CardsAmount uint   `json:"cards_amount"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
