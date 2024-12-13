package entities

type Producer struct {
	ID           int         `json:"id" db:"id"`
	Municipality string      `json:"municipality" db:"municipality"`
	Fccw         string      `json:"fccw" db:"fccw"`
	HazardClass  string      `json:"hazard_class" db:"hazard_class"`
	Organization string      `json:"organization" db:"organization"`
	WasteType    []WasteType `json:"waste_type" db:"waste_type"`
}
