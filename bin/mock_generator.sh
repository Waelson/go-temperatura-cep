mockgen -source=../internal/service/application.go -destination=../internal/tests/service/mock_application.go ApplicationService
mockgen -source=../internal/service/integration.go -destination=../internal/tests/service/mock_integration.go IntegrationService
mockgen -source=../internal/requester/requester.go -destination=../internal/tests/requester/mock_requester.go HttpRequest
mockgen -source=../internal/controller/application.go -destination=../internal/tests/controller/mock_application.go ApplicationController