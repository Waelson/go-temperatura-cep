package service

import (
	"encoding/json"
	"fmt"
	"github.com/Waelson/go-temperatura-cep/internal/model"
	"github.com/Waelson/go-temperatura-cep/internal/requester"
	"net/http"
)

type IntegrationService interface {
	GetCep(cep string) (model.CepResponse, error)
	GetTemperatura(cidade string) (model.TemperaturaResponse, error)
}

type integrationService struct {
	httpRequest requester.HttpRequest
	urls        model.Url
}

func (s *integrationService) GetCep(cep string) (model.CepResponse, error) {
	fmt.Println(fmt.Sprintf("Pesquisando CEP %s", cep))
	res, status, err := s.httpRequest.MakeRequest(s.urls.GetCep(cep))

	if status == http.StatusBadRequest {
		return model.CepResponse{}, model.InvalidCepError
	}

	if status == http.StatusNotFound {
		return model.CepResponse{}, model.CepNotFoundError
	}

	if status != http.StatusOK {
		return model.CepResponse{}, model.InternalError
	}

	var cepResponse model.CepResponse
	if err = json.Unmarshal([]byte(res), &cepResponse); err != nil {
		return model.CepResponse{}, err
	}

	if cepResponse.Error {
		return model.CepResponse{}, model.CepNotFoundError
	}

	return cepResponse, nil
}

func (s *integrationService) GetTemperatura(cidade string) (model.TemperaturaResponse, error) {
	fmt.Println(fmt.Sprintf("Pesquisando Temperatura %s", cidade))
	res, status, err := s.httpRequest.MakeRequest(s.urls.GetTemperatura(cidade))

	if status != http.StatusOK {
		return model.TemperaturaResponse{}, model.InternalError
	}

	var temperaturaResponse model.TemperaturaResponse
	if err = json.Unmarshal([]byte(res), &temperaturaResponse); err != nil {
		return model.TemperaturaResponse{}, err
	}
	return temperaturaResponse, nil
}

func NewIntegrationService(httpRequest requester.HttpRequest, urls model.Url) IntegrationService {
	return &integrationService{
		httpRequest: httpRequest,
		urls:        urls,
	}
}
