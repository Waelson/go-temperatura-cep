package service

import (
	"fmt"
	"github.com/Waelson/go-temperatura-cep/internal/model"
	"strconv"
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
	cep = strings.TrimSpace(cep)
	if cep == "" || len(cep) != 8 || !isNumber(cep) {
		return model.ApplicationResponse{}, model.InvalidCepError
	}

	cepResponse, err := s.integrationService.GetCep(cep)
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

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func NewApplicationService(service IntegrationService) ApplicationService {
	return &applicationService{integrationService: service}
}
