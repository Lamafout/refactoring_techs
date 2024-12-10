package presentation

import (
	"encoding/json"
	"net/http"
	"refactoring_tech/presentation/controllers"
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
