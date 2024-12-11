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

func (r *RepositoryImpl) GetListOfTechs() (*[]entities.TechShort, error) {
	query := `
	SELECT 
			t.id, t.name, t.specs, t.perfomance, 
			a.id AS assignment_id, a.name AS assignment_name,
			u.id AS use_cases_id, u.name AS use_cases_name,
			e.id AS expert_info_id, e.authority_name, e.date, e.conclusion
	FROM 
			techs t
	LEFT JOIN assignments a ON t.assignment = a.id
	LEFT JOIN use_cases u ON t.use_case = u.id
	LEFT JOIN expert_info e ON t.expert_info = e.id`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var techs []entities.TechShort
	for rows.Next() {
		var tech entities.Tech
		var assignment entities.Assignment
		var useCases entities.UseCases
		var expertInfo entities.ExpertInfo

		err := rows.Scan(
			&tech.ID, &tech.Name, &tech.Specs, &tech.Perfomance,
			&assignment.ID, &assignment.Name,
			&useCases.ID, &useCases.Name,
			&expertInfo.ID, &expertInfo.AuthorityNameCharacter, &expertInfo.Date, &expertInfo.Conclusion,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		tech.Assignment = assignment
		tech.UseCases = useCases
		tech.ExpertInfo = expertInfo

		techs = append(techs, *entities.NewTechShortFromTech(tech))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &techs, nil
}

func (r *RepositoryImpl) GetConcreteTech(id int) (*entities.Tech, error) {
	query := `
  SELECT 
    t.id AS tech_id,
    t.name AS tech_name,
    t.assignment,
    a.name AS assignment_name,
    t.specs,
    t.resources,
    r.energy,
    r.water,
    r.neutralization_and_disposal,
    t.performance,
    sw.id AS secondary_waste_id,
    sw.mass,
    sw.volume,
	f.code AS fccw_code,
    f.name AS fccw_name,
    cpta.id AS cpta_id,
    cpta.name AS cpta_name,
	t.contacts,
	t.user_contacts
  FROM 
    techs t
  LEFT JOIN assignments a ON t.assignment = a.id
  LEFT JOIN resources r ON t.resources = r.id
  LEFT JOIN secondary_waste sw ON sw.id IN (
    SELECT sec 
    FROM sec_in_tech 
    WHERE tech = t.id
  )
  LEFT JOIN fccw f ON sw.fccw = f.id
  LEFT JOIN cpta_in_tech cit ON cit.tech = t.id
  LEFT JOIN cpta cpta ON cit.cpta = cpta.id
  LEFT JOIN contacts dev_contacts ON t.contacts = dev_contacts.id
  LEFT JOIN contacts user_contacts ON t.user_contacts = user_contacts.id
  WHERE 
    t.id = $1;
  `

	rows, err := r.db.Query(query, id)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var tech entities.Tech
	var secondaryWastes []entities.SecondaryWaste
	var cptas []entities.Cpta

	for rows.Next() {
		var assignment entities.Assignment
		var resources entities.Resources
		var secondaryWaste entities.SecondaryWaste
		var cpta entities.Cpta
		var fccwCode sql.NullString
		var fccwName sql.NullString
		var devContacts sql.NullInt64
		var userContacts sql.NullInt64

		err := rows.Scan(
			&tech.ID, &tech.Name, &tech.Assignment.ID, &assignment.Name,
			&tech.Specs, &tech.Resources.ID,
			&resources.Energy, &resources.Water, &resources.NeutralizationAndDisposal, &tech.Perfomance,
			&secondaryWaste.ID, &secondaryWaste.Mass, &secondaryWaste.Volume,
			&fccwCode, &fccwName,
			&cpta.ID, &cpta.Name,
			&devContacts, &userContacts,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		tech.Assignment = assignment
		tech.Resources = resources

		if secondaryWaste.ID != 0 {
			if fccwName.Valid {
				secondaryWaste.Name = fccwName.String
			}
			secondaryWastes = append(secondaryWastes, secondaryWaste)
		}

		if cpta.ID != 0 {
			cptas = append(cptas, cpta)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	tech.SecondaryWaste = secondaryWastes
	tech.Cpta = cptas

	return &tech, nil
}

func ConvertTechToModel(tech entities.Tech) models.TechModel {
	return models.TechModel{
		ID:           tech.ID,
		Name:         tech.Name,
		Assignment:   tech.Assignment.ID,
		Specs:        tech.Specs,
		Resources:    tech.Resources.ID,
		Perfomance:   tech.Perfomance,
		Contacts:     tech.Contacts.ID,
		UserContacts: tech.UserContacts.ID,
		Fccw:         tech.Fccw.ID,
		UseCases:     tech.UseCases.ID,
		ExpertInfo:   tech.ExpertInfo.ID,
	}
}

func ConvertContactsToModel(contacts entities.Contacts) models.ContactsModel {
	return models.ContactsModel{
		ID:      contacts.ID,
		Address: contacts.Address,
		Fax:     contacts.Fax,
		Phone:   contacts.Phone,
		Site:    contacts.Site,
	}
}

func ConvertResourcesToModel(resources entities.Resources) models.ResourcesModel {
	return models.ResourcesModel{
		ID:                        resources.ID,
		Energy:                    resources.Energy,
		Water:                     resources.Water,
		NeutralizationAndDisposal: resources.NeutralizationAndDisposal,
	}
}

func ConvertUseCasesToModel(useCases entities.UseCases) models.UseCasesModel {
	return models.UseCasesModel{
		ID:   useCases.ID,
		Name: useCases.Name,
	}
}

func ConvertExpertInfoToModel(expertInfo entities.ExpertInfo) models.ExpertInfoModel {
	return models.ExpertInfoModel{
		ID:                     expertInfo.ID,
		AuthorityNameCharacter: expertInfo.AuthorityNameCharacter,
		Date:                   expertInfo.Date,
		Conclusion:             expertInfo.Conclusion,
	}
}

func ConvertAssigmentToModel(assignment entities.Assignment) models.AssignmentsModel {
	return models.AssignmentsModel{
		ID:   assignment.ID,
		Name: assignment.Name,
	}
}

func ConvertToSecTechModel(sec entities.SecondaryWaste, tech entities.Tech) models.SecInTechModel {
	return models.SecInTechModel{
		Sec:  sec.ID,
		Tech: tech.ID,
	}
}

func ConvertSecondaryWasteToModel(sec entities.SecondaryWaste) models.SecondaryWasteModel {
	return models.SecondaryWasteModel{
		ID:     sec.ID,
		Mass:   sec.Mass,
		Volume: sec.Volume,
		Fccw:   sec.FccwId,
	}
}

func ConvertFccwToModel(fccw entities.Fccw) models.FccwModel {
	return models.FccwModel{
		ID:   fccw.ID,
		Name: fccw.Name,
		Code: fccw.Code,
	}
}

func ConvertCptaToModel(cpta entities.Cpta) models.CptaModel {
	return models.CptaModel{
		ID:   cpta.ID,
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
