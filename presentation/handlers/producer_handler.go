package handlers

import (
	_ "encoding/json"
	"io"
	"net/http"
	"refactoring_tech/domain/entities"
	"refactoring_tech/presentation/controllers"
	"refactoring_tech/presentation/requests"
	_ "strconv"
)

type ProducerHandler struct {
	Controller controllers.ProducerController
}

func NewProducerHandler(producerController controllers.ProducerController) *ProducerHandler { //initialize the handler
	return &ProducerHandler{
		Controller: producerController,
	}
}

func (h *ProducerHandler) ProducerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			h.CreateProducerHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProducerHandler) CreateProducerHandler(w http.ResponseWriter, r *http.Request) {
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

	err = h.Controller.InsertProducer(producer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
