package entities

type TechShort struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Assignment Assignment `json:"assignment"`
	Specs      string     `json:"specs"`
	UseCases   UseCases   `json:"useCases"`
	ExpertInfo ExpertInfo `json:"expertInfo"`
	Fccw       []Fccw     `json:"fccw"`
}

func NewTechShortFromTech(tech Tech) *TechShort {
	return &TechShort{
		ID:         tech.ID,
		Name:       tech.Name,
		Assignment: tech.Assignment,
		Specs:      tech.Specs,
		UseCases:   tech.UseCases,
		ExpertInfo: tech.ExpertInfo,
		Fccw:       tech.Fccw,
	}
}
