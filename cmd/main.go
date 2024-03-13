package main

import (
	"fmt"
	"github.com/Waelson/go-temperatura-cep/internal/controller"
	"github.com/Waelson/go-temperatura-cep/internal/model"
	"github.com/Waelson/go-temperatura-cep/internal/requester"
	"github.com/Waelson/go-temperatura-cep/internal/service"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Iniciando aplicacao")

	httpRequest := requester.NewHttpRequest()
	urls := model.NewModel()
	integrationService := service.NewIntegrationService(httpRequest, urls)
	applicationService := service.NewApplicationService(integrationService)
	applicationController := controller.NewApplicationController(applicationService)

	http.Handle("/", http.HandlerFunc(applicationController.Handler))

	log.Println("Iniciando o servidor na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %s\n", err)
	}

}
