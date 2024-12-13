package models

type WasteTypeModel struct {
	ID   int    `json:"id" db:"id"`
	Type string `json:"type" db:"type"`
}
