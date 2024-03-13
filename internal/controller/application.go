package controller

import (
	"github.com/Waelson/go-temperatura-cep/internal/service"
	"net/http"
)

type ApplicationController interface {
	Handler(w http.ResponseWriter, r *http.Request)
}
type applicationController struct {
	service service.ApplicationService
}

func (c *applicationController) Handler(w http.ResponseWriter, r *http.Request) {

}

func NewApplicationController(service service.ApplicationService) ApplicationController {
	return &applicationController{service: service}
}
