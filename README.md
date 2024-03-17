### Projeto
Projeto para consulta de temperatura baseado nno CEP

### Pré Requisitos
Golang 1.22.1

### Como executar?
```
docker-compose up --build
```
Acesse a URL abaixo
```
http://localhost:8080?cep=<cep>
```
### Cenários de Teste
Sucesso (200)
```
http://localhost:8080?cep=64770000
```
Cep inválido (422)
```
http://localhost:8080?cep=647700
```
Cep não encontrado (404)
```
http://localhost:8080?cep=00000000
```
### Google Cloud Run
Endereço de teste
```
https://go-temperatura-cep-g4wj6iiisa-uc.a.run.app/?cep=64770000
```