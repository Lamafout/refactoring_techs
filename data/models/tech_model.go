package models

type TechModel struct {
	ID int 		`json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Assignment int `json:"assignment" db:"assignment"`
	Specs string `json:"specs" db:"specs"`
	Resources int `json:"resources" db:"resources"`
	Perfomance string `json:"perfomance" db:"perfomance"`
	Contacts int `json:"contacts" db:"contacts"`
	UseCases int `json:"use_cases_case" db:"use_cases_case"`
	ExpertInfo int `json:"expert_info" db:"expert_info"`
}
