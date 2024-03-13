package service

import (
	"fmt"
	"github.com/Waelson/go-temperatura-cep/internal/model"
	"strings"
)

type ApplicationService interface {
	GetTemperature(cep string) (model.ApplicationResponse, error)
}

type applicationService struct {
	integrationService IntegrationService
}

func (s *applicationService) GetTemperature(cep string) (model.ApplicationResponse, error) {
	fmt.Println(fmt.Sprintf("Vamos pesquisar o CEP %s", cep))

	if strings.TrimSpace(cep) == "" {
		return model.ApplicationResponse{}, model.InvalidCepError
	}

	cepResponse, err := s.integrationService.GetCep(strings.TrimSpace(cep))
	if err != nil {
		return model.ApplicationResponse{}, err
	}

	temperaturaResponse, err := s.integrationService.GetTemperatura(strings.TrimSpace(cepResponse.Localidade))
	if err != nil {
		return model.ApplicationResponse{}, err
	}

	tempF := temperaturaResponse.Current.TempC*1.8 + 32
	tempK := temperaturaResponse.Current.TempC + 273

	return model.ApplicationResponse{
		TempC: temperaturaResponse.Current.TempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}

func NewApplicationService(service IntegrationService) ApplicationService {
	return &applicationService{integrationService: service}
}
