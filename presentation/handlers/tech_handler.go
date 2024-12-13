package handlers

import (
	"encoding/json"
	"net/http"
	"refactoring_tech/presentation/controllers"
	"strconv"
)

type TechHandler struct {
	Controller controllers.TechsController
}

func NewTechHandler(techsController controllers.TechsController) *TechHandler { //initialize the handler
	return &TechHandler{
		Controller: techsController,
	}
}

func (h *TechHandler) TechHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.GetConcreteTechHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TechHandler) GetTechsHandler(w http.ResponseWriter, r *http.Request) {
	techs, err := h.Controller.GetTechs()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(techs)
}

func (h *TechHandler) GetConcreteTechHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tech, err := h.Controller.GetTechById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tech)
}
