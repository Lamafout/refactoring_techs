package entities

type Tech struct {
	ID             int              `json:"id" bd:"id"`
	Name           string           `json:"name" bd:"name"`
	Assignment     Assignment       `json:"assignment" bd:"assignment"`
	Specs          string           `json:"specs" bd:"specs"`
	Resources      Resources        `json:"resources" bd:"resources"`
	Perfomance     string           `json:"perfomance" bd:"perfomance"`
	SecondaryWaste []SecondaryWaste `json:"secondary_waste" bd:"secondary_waste"`
	Contacts       Contacts         `json:"contacts" bd:"contacts"`
	UseCases       UseCases         `json:"use_cases_case" bd:"use_cases_case"`
	ExpertInfo     ExpertInfo       `json:"expert_info" bd:"expert_info"`
	Fccw           []Fccw           `json:"fccw" bd:"fccw"`
	UserContacts   []Contacts       `json:"user_contacts" bd:"user_contacts"`
	Cpta           []Cpta           `json:"cpta" bd:"cpta"`
}
