package service

import "github.com/Waelson/go-temperatura-cep/internal/model"

type ApplicationService interface {
	GetTemperature(cep string) (model.ApplicationResponse, error)
}

type applicationService struct {
	integrationService IntegrationService
}

func (s *applicationService) GetTemperature(cep string) (model.ApplicationResponse, error) {
	return model.ApplicationResponse{}, nil
}

func NewApplicationService(service IntegrationService) ApplicationService {
	return &applicationService{integrationService: service}
}
