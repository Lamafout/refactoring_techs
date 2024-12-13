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
			t.id, t.name, t.specs, t.performance, 
			a.id AS assignment_id, a.name AS assignment_name,
			u.id AS use_cases_id, u.name AS use_cases_name,
			e.id AS expert_info_id, e.authority_name, e.date, e.conclusion,
			f.id, f.code, f.name
	FROM 
			techs t
	LEFT JOIN assignments a ON t.assignment = a.id
	LEFT JOIN use_cases u ON t.use_case = u.id
	LEFT JOIN expert_info e ON t.expert_info = e.id
	LEFT JOIN fccw_in_tech fit ON t.id = fit.tech
	LEFT JOIN fccw f ON fit.fccw = f.id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var techs []entities.TechShort
	var techMap = make(map[int]*entities.Tech) // Карта для хранения временных объектов Tech
	for rows.Next() {
		var tech entities.Tech
		var assignment entities.Assignment
		var useCases entities.UseCases
		var expertInfo entities.ExpertInfo
		var fccw entities.Fccw

		err := rows.Scan(
			&tech.ID, &tech.Name, &tech.Specs, &tech.Perfomance,
			&assignment.ID, &assignment.Name,
			&useCases.ID, &useCases.Name,
			&expertInfo.ID, &expertInfo.AuthorityNameCharacter, &expertInfo.Date, &expertInfo.Conclusion,
			&fccw.ID, &fccw.Code, &fccw.Name,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Если в карте еще нет такого tech, добавляем его
		if _, exists := techMap[tech.ID]; !exists {
			tech.Assignment = assignment
			tech.UseCases = useCases
			tech.ExpertInfo = expertInfo
			tech.Fccw = []entities.Fccw{} // Инициализация массива для Fccw
			techMap[tech.ID] = &tech
		}

		// Добавляем fccw в массив для tech
		techMap[tech.ID].Fccw = append(techMap[tech.ID].Fccw, fccw)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Переводим карту в слайс
	for _, tech := range techMap {
		techs = append(techs, *entities.NewTechShortFromTech(*tech))
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
	f_sec.code AS sec_fccw_code,
    f_sec.name AS sec_fccw_name,
	f.code AS fccw_code,
    f.name AS fccw_name,
    cpta.id AS cpta_id,
	cpta.code,
    cpta.name,
	dev_contacts.address AS dc_address,
	dev_contacts.phone AS dc_phone,
	dev_contacts.fax AS dc_fax,
	dev_contacts.site AS dc_site,
	user_contacts.address AS uc_address,
	user_contacts.phone AS uc_phone,
	user_contacts.fax AS uc_fax,
	user_contacts.site AS uc_site,
	usca.id AS usca_id,
	usca.name AS usca_name,
	ei.id AS ei_id, ei.authority_name, ei.date, ei.conclusion
  FROM 
    techs t
  LEFT JOIN assignments a ON t.assignment = a.id
  LEFT JOIN resources r ON t.resources = r.id
  LEFT JOIN secondary_waste sw ON sw.id IN (
    SELECT sec 
    FROM sec_in_tech 
    WHERE tech = t.id
  )
  LEFT JOIN fccw_in_tech fit ON fit.tech = t.id
  LEFT JOIN fccw f ON fit.fccw = f.id
  LEFT JOIN fccw f_sec ON f_sec.id = sw.fccw
  LEFT JOIN cpta_in_tech cit ON cit.tech = t.id
  LEFT JOIN cpta cpta ON cit.cpta = cpta.id
  LEFT JOIN contacts dev_contacts ON t.contacts = dev_contacts.id
  LEFT JOIN user_in_tech uit ON uit.tech = t.id
  LEFT JOIN contacts user_contacts ON uit.user = user_contacts.id
  LEFT JOIN use_cases usca ON usca.id = t.use_case
  LEFT JOIN expert_info ei ON t.expert_info = ei.id
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
	var fccws []entities.Fccw
	var userContactsArray []entities.Contacts

	for rows.Next() {
		var assignment entities.Assignment
		var resources entities.Resources
		var secondaryWaste entities.SecondaryWaste
		var cpta entities.Cpta
		var fccw entities.Fccw
		var devContacts entities.Contacts
		var userContacts entities.Contacts
		var useCases entities.UseCases
		var expertInfo entities.ExpertInfo

		err := rows.Scan(
			&tech.ID, &tech.Name, 
			&assignment.ID, &assignment.Name,
			&tech.Specs, 
			&resources.ID, &resources.Energy, &resources.Water, &resources.NeutralizationAndDisposal, 
			&tech.Perfomance,
			&secondaryWaste.ID, &secondaryWaste.Mass, &secondaryWaste.Volume, &secondaryWaste.Code, &secondaryWaste.Name,
			&fccw.Code, &fccw.Name,
			&cpta.ID, &cpta.Code, &cpta.Name, 
			&devContacts.Address, &devContacts.Phone, &devContacts.Fax, &devContacts.Site,
			&userContacts.Address, &userContacts.Phone, &userContacts.Fax, &userContacts.Site,
			&useCases.ID, &useCases.Name,
			&expertInfo.ID, &expertInfo.AuthorityNameCharacter, &expertInfo.Date, &expertInfo.Conclusion,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}


		tech.Assignment = assignment
		tech.Resources = resources
		tech.Contacts = devContacts
		tech.UserContacts = userContactsArray
		tech.UseCases = useCases
		tech.ExpertInfo = expertInfo

		if secondaryWaste.ID != 0 {
			secondaryWastes = append(secondaryWastes, secondaryWaste)
		}

		cptas = append(cptas, cpta)
		fccws = append(fccws, fccw)
		userContactsArray = append(userContactsArray, userContacts)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	tech.SecondaryWaste = secondaryWastes
	tech.Cpta = cptas
	tech.Fccw = fccws
	tech.UserContacts = userContactsArray

	return &tech, nil
}

func (r *RepositoryImpl) InsertProducer(producer entities.Producer) error {
	var fccwID int
	err := r.db.QueryRow(`SELECT id FROM fccw WHERE code = $1`, producer.Fccw).Scan(&fccwID)
	if err != nil {
		log.Fatalf("Failed to find fccw id: %v", err)
	}

	var producerID int
	err = r.db.QueryRow(`
		INSERT INTO producers (municipality, fccw, hazard_class, organization)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		producer.Municipality, fccwID, producer.HazardClass, producer.Organization).Scan(&producerID)
	if err != nil {
		log.Fatalf("Erorr: %v", err)
	}

	for _, wasteType := range producer.WasteType {
		var wasteTypeID int
		err := r.db.QueryRow(`
			INSERT INTO waste_types (type) 
			VALUES ($1) ON CONFLICT (type) DO UPDATE SET type = EXCLUDED.type
			RETURNING id`, wasteType.Type).Scan(&wasteTypeID)
		if err != nil {
			log.Fatalf("Erorr: %v", err)
		}

		_, err = r.db.Exec(`
			INSERT INTO waste_types_in_producers (producer_id, waste_type_id) 
			VALUES ($1, $2) ON CONFLICT DO NOTHING`,
			producerID, wasteTypeID)
		if err != nil {
			log.Fatalf("Failed to insert new producer: %v", err)
		}
	}

	return nil
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
