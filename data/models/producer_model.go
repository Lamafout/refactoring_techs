package models

type ProducerModel struct {
	ID           int    `json:"id" db:"id"`
	Municipality string `json:"municipality" db:"municipality"`
	Fccw         int    `json:"fccw" db:"fccw"`
	HazardClass  string `json:"hazard_class" db:"hazard_class"`
	Organization string `json:"organization" db:"organization"`
	WasteType    int    `json:"waste_type" db:"waste_type"`
}
