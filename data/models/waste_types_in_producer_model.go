package models

type WasteTypeInProducer struct {
	WasteType int `json:"wasteType" db:"waste_type"`
	Producer  int `json:"producer" db:"producer"`
}
