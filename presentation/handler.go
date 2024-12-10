package presentation

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	TechsController TechsController
}

func NewHandler(techsController TechsController) *Handler { //initialize the handler
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
