package service

import (
	"github.com/Waelson/go-temperatura-cep/internal/model"
	"github.com/Waelson/go-temperatura-cep/internal/requester"
)

type IntegrationService interface {
	GetCep(cep string) (model.CepResponse, error)
	GetTemperatura(cidade string) (model.TemperaturaResponse, error)
}

type integrationService struct {
	httpRequest requester.HttpRequest
}

func (s *integrationService) GetCep(cep string) (model.CepResponse, error) {
	return model.CepResponse{}, nil
}

func (s *integrationService) GetTemperatura(cidade string) (model.TemperaturaResponse, error) {
	return model.TemperaturaResponse{}, nil
}

func NewIntegrationService(httpRequest requester.HttpRequest) IntegrationService {
	return &integrationService{
		httpRequest: httpRequest,
	}
}
