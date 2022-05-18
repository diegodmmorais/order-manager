# order-manager

## Clean architecture
Arquitetura Limpa (Clean Architecture) tem como objetivo de promover a implementação de sistemas que favorecem:
    1 - reusabilidade de código
    2 - coesão 
    3 - independência de tecnologia
    4 - testabilidade.

## Comando para executar coverage
```shell
go test ./... --coverprofile=cover.out && go tool cover --html=cover.out
``` 
## Comando para compilar
```shell
go build
``` 

## Comando para instalar o mondodb e postgres
```shell
docker-compose up -d
``` 
Link para acessar mongo-express: http://localhost:8081/

link para acessar pgadmin4: http://localhost:8082/
* Usuario: admin@admin.com
* Senha: admin123
