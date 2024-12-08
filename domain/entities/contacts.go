package entities

type Contacts struct {
	ID int `json:"id" db:"id"`
	Address string `json:"address" db:"address"`
	Phone string `json:"phone" db:"phone"`
	Fax string `json:"fax" db:"fax"`
	Site string `json:"site" db:"site"`
}