package entities

type SecondaryWaste struct {
	ID     int    `json:"id" db:"id"`
	Mass   int    `json:"mass" db:"mass"`
	Volume int    `json:"volume" db:"volume"`
	FccwId int    `json:"fccw_id" db:"fccw_id"`
	Code   string `json:"code" db:"code"` // code from fccw
	Name   string `json:"name" db:"name"` // name from fccw
}
