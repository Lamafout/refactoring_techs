package models

type SecondaryWasteModel struct {
	ID int		`json:"id" db:"id"`
	Mass int 	`json:"mass" db:"mass"`
	Volume int  `json:"volume" db:"volume"`
	Fccw int 	`json:"fccw" db:"fccw"`
}