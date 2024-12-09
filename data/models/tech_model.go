package models

type TechModel struct {
	ID int 		`json:"id" bd:"id"`
	Name string `json:"name" bd:"name"`
	Assignment int `json:"assignment" bd:"assignment"`
	Specs string `json:"specs" bd:"specs"`
	Resources int `json:"resources" bd:"resources"`
	Perfomance string `json:"perfomance" bd:"perfomance"`
	Contacts int `json:"contacts" bd:"contacts"`
	UseCases int `json:"use_cases_case" bd:"use_cases_case"`
	ExpertInfo int `json:"expert_info" bd:"expert_info"`
	Fccw int `json:"fccw" bd:"fccw"`
	UserContacts int `json:"user_contacts" bd:"user_contacts"`
}