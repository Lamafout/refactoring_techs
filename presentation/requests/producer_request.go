package requests

import (
	"encoding/json"
	"refactoring_tech/domain/entities"
)

type ProducerRequest struct {
	Municipality string          `json:"municipality"`
	FkkoCode     string          `json:"fkkoCode"`
	WasteTypes   []WasteTypeJSON `json:"wasteTypes"`
	HazardClass  string          `json:"hazardClass"`
	Organization string          `json:"organization"`
}

type WasteTypeJSON struct {
	Type string `json:"type"`
}

func ConvertClientRequestToProducer(requestBody []byte) (entities.Producer, error) {
	var clientRequest ProducerRequest
	err := json.Unmarshal(requestBody, &clientRequest)
	if err != nil {
		return entities.Producer{}, err
	}

	wasteTypes := make([]entities.WasteType, len(clientRequest.WasteTypes))
	for i, wt := range clientRequest.WasteTypes {
		wasteTypes[i] = entities.WasteType{
			Type: wt.Type,
		}
	}

	producer := entities.Producer{
		Municipality: clientRequest.Municipality,
		Fccw:         clientRequest.FkkoCode,
		HazardClass:  clientRequest.HazardClass,
		Organization: clientRequest.Organization,
		WasteType:    wasteTypes,
	}

	return producer, nil
}