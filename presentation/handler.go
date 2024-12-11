package presentation

import (
	"encoding/json"
	"net/http"
	"refactoring_tech/presentation/controllers"
	"strconv"
)

type Handler struct {
	TechsController controllers.TechsController
}

func NewHandler(techsController controllers.TechsController) *Handler { //initialize the handler
	return &Handler{
		TechsController: techsController,
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