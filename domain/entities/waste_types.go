package entities

type WasteType struct {
	ID   int    `json:"id" db:"id"`
	Type string `json:"type" db:"type"`
}
