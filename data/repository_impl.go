package db

import (
	"database/sql"
	"log"
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
	query := `
	SELECT 
			t.id, t.name, t.specs, t.perfomance, 
			a.id AS assignment_id, a.name AS assignment_name,
			r.id AS resources_id, r.energy, r.water, r.neutralization_and_disposal,
			c.id AS contacts_id, c.address, c.phone, c.fax, c.site,
			u.id AS use_cases_id, u.name AS use_cases_name,
			e.id AS expert_info_id, e.authority_name, e.date, e.conclusion,
			f.id AS fccw_id, f.name AS fccw_name, f.code AS fccw_code,
			cw.id AS secondary_waste_id, cw.mass, cw.volume, cw.fccw AS secondary_waste_fccw_id
	FROM 
			techs t
	LEFT JOIN assignments a ON t.assignment = a.id
	LEFT JOIN resources r ON t.resources = r.id
	LEFT JOIN contacts c ON t.contacts = c.id
	LEFT JOIN use_cases u ON t.use_case = u.id
	LEFT JOIN expert_info e ON t.expert_info = e.id
	LEFT JOIN fccw f ON t.fccw = f.id
	LEFT JOIN secondary_waste cw ON t.secondary_waste = cw.id`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var techs []entities.Tech
	for rows.Next() {
		var tech entities.Tech
		var assignment entities.Assignment
		var resources entities.Resources
		var contacts entities.Contacts
		var useCases entities.UseCases
		var expertInfo entities.ExpertInfo
		var fccw entities.Fccw
		var secondaryWaste entities.SecondaryWaste

		err := rows.Scan(
			&tech.ID, &tech.Name, &tech.Specs, &tech.Perfomance,
			&assignment.ID, &assignment.Name,
			&resources.ID, &resources.Energy, &resources.Water, &resources.NeutralizationAndDisposal,
			&contacts.ID, &contacts.Address, &contacts.Phone, &contacts.Fax, &contacts.Site,
			&useCases.ID, &useCases.Name,
			&expertInfo.ID, &expertInfo.AuthorityNameCharacter, &expertInfo.Date, &expertInfo.Conclusion,
			&fccw.ID, &fccw.Name, &fccw.Code,
			&secondaryWaste.ID, &secondaryWaste.Mass, &secondaryWaste.Volume, &secondaryWaste.FccwId,
		)
		if err != nil {
			return nil, err
		}

		tech.Assignment = assignment
		tech.Resources = resources
		tech.Contacts = contacts
		tech.UseCases = useCases
		tech.ExpertInfo = expertInfo
		tech.Fccw = fccw
		tech.SecondaryWaste = secondaryWaste

		techs = append(techs, tech)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &techs, nil
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