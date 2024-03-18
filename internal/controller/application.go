package controller

import (
	"encoding/json"
	"errors"
	"github.com/Waelson/go-temperatura-cep/internal/model"
	"github.com/Waelson/go-temperatura-cep/internal/service"
	"net/http"
	"strings"
)

type ApplicationController interface {
	Handler(w http.ResponseWriter, r *http.Request)
}
type applicationController struct {
	service service.ApplicationService
}

func (c *applicationController) Handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	cep := queryParams.Get("cep")

	response, err := c.service.GetTemperature(strings.TrimSpace(cep))

	w.Header().Set("Content-Type", "application/json")

	if errors.Is(err, model.InvalidCepError) {
		//w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
	} else if errors.Is(err, model.CepNotFoundError) {
		//w.WriteHeader(http.StatusNotFound)
		http.Error(w, "can not find zipcode", http.StatusNotFound)
	} else if errors.Is(err, model.InternalError) {
		//w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "internal error", http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func NewApplicationController(service service.ApplicationService) ApplicationController {
	return &applicationController{service: service}
}
