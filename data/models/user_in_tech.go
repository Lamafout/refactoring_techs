package models

type UserInTech struct {
	User int `json:"user" db:"user"`
	Tech int `json:"tech" db:"tech"`
}
