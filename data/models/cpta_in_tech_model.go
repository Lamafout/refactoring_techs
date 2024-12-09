package models

type CptaInTechModel struct {
	Cpta int `json:"cpta" db:"cpta"`
	Tech int `json:"tech" db:"tech"`
}