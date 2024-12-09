package db

import (
	"database/sql"
	"refactoring_tech/data/models"
	"refactoring_tech/domain/entities"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) GetListOfTechs() (*[]entities.Tech, error) {
	// TODO add query for get fulll data about techs
	return nil, nil
}

func ConvertTechToModel(tech entities.Tech) models.TechModel {
	return models.TechModel{
		ID: tech.ID,
		Name: tech.Name,
		Assignment: tech.Assignment.ID,
		Specs: tech.Specs,
		Resources: tech.Resources.ID,
		Perfomance: tech.Perfomance,
		Contacts: tech.Contacts.ID,
		UserContacts: tech.UserContacts.ID,
		Fccw: tech.Fccw.ID,
		UseCases: tech.UseCases.ID,
		ExpertInfo: tech.ExpertInfo.ID,
	}
}

func ConvertContactsToModel(contacts entities.Contacts) models.ContactsModel {
	return models.ContactsModel{
		ID: contacts.ID,
		Address: contacts.Address,
		Fax: contacts.Fax,
		Phone: contacts.Phone,
		Site: contacts.Site,
	}
}

func ConvertResourcesToModel(resources entities.Resources) models.ResourcesModel {
	return models.ResourcesModel{
		ID: resources.ID,
		Energy: resources.Energy,
		Water: resources.Water,
		NeutralizationAndDisposal: resources.NeutralizationAndDisposal,
	}
}

func ConvertUseCasesToModel(useCases entities.UseCases) models.UseCasesModel {
	return models.UseCasesModel{
		ID: useCases.ID,
		Name: useCases.Name,
	}
}

func ConvertExpertInfoToModel(expertInfo entities.ExpertInfo) models.ExpertInfoModel {
	return models.ExpertInfoModel{
		ID: expertInfo.ID,
		AuthorityNameCharacter: expertInfo.AuthorityNameCharacter,
		Date: expertInfo.Date,
		Conclusion: expertInfo.Conclusion,
	}
}

func ConvertAssigmentToModel(assignment entities.Assignment) models.AssignmentsModel {
	return models.AssignmentsModel{
		ID: assignment.ID,
		Name: assignment.Name,
	}
}

func ConvertToSecTechModel(sec entities.SecondaryWaste, tech entities.Tech) models.SecInTechModel {
	return models.SecInTechModel{
		Sec: sec.ID,
		Tech: tech.ID,
	}
}

func ConvertSecondaryWasteToModel(sec entities.SecondaryWaste) models.SecondaryWasteModel {
	return models.SecondaryWasteModel{
		ID: sec.ID,
		Mass: sec.Mass,
		Volume: sec.Volume,
		Fccw: sec.FccwId,
	}
}

func ConvertFccwToModel(fccw entities.Fccw) models.FccwModel {
	return models.FccwModel{
		ID: fccw.ID,
		Name: fccw.Name,
		Code: fccw.Code,
	}
}

func ConvertCptaToModel(cpta entities.Cpta) models.CptaModel {
	return models.CptaModel{
		ID: cpta.ID,
		Name: cpta.Name,
		Code: cpta.Code,
	}
}

func ConvertToCptanTechModel(cpta entities.Cpta, tech entities.Tech) models.CptaInTechModel {
	return models.CptaInTechModel{
		Cpta: cpta.ID,
		Tech: tech.ID,
	}
}