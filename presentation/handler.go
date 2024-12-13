package presentation

import (
	"encoding/json"
	"io"
	"net/http"
	"refactoring_tech/domain/entities"
	"refactoring_tech/presentation/controllers"
	"refactoring_tech/presentation/requests"
	"strconv"
)

type Handler struct {
	TechsController controllers.TechsController
	ProducerController controllers.ProducersController
}

func NewHandler(techsController controllers.TechsController, producerController controllers.ProducersController) *Handler { //initialize the handler
	return &Handler{
		TechsController: techsController,
		ProducerController: producerController,
	}
}

func (h *Handler) GetTechsHandler(w http.ResponseWriter, r *http.Request) {
	techs, err := h.TechsController.GetTechs()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(techs) 
}

func (h *Handler) GetConcreteTechHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tech, err := h.TechsController.GetTechById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tech)
}

func (h *Handler) CreateProducerHandler(w http.ResponseWriter, r *http.Request) {
	var producer entities.Producer
	requstBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	producer, err = requests.ConvertClientRequestToProducer(requstBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.ProducerController.InsertProducer(producer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 
	w.WriteHeader(http.StatusCreated)
}