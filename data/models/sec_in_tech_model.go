package models

type SecInTechModel struct {
	Sec int 	`json:"sec" db:"sec"`
	Tech int 	`json:"tech" db:"tech"`
}