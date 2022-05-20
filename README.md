# order-manager
![alt text](https://github.com/diegodmmorais/order-manager/blob/main/golang-with-clean-architecture.drawio.png)
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


## Script postgresql
```sql

-- Database: labsit

-- DROP DATABASE IF EXISTS labsit;

CREATE DATABASE labsit
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

------------------------------------------------------------------

-- Table: public.orders

-- DROP TABLE IF EXISTS public.orders;

CREATE TABLE IF NOT EXISTS public.orders
(
    id character varying COLLATE pg_catalog."default" NOT NULL,
    data text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT orders_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.orders
    OWNER to postgres;

```

## Postman Payload Test
```json
curl --location --request POST 'http://localhost:1323/api/v1/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "customer_id": "#1",
    "shipping_address": {
        "street": "Rua new street",
        "number": "777",
        "zip_code": "98897-890",
        "city": "São paulo - SP"
    },
    "freight": 55.98,
    "items": [
        {
            "product_id": "#1",
            "amount": 1
        },
        {
            "product_id": "#1",
            "amount": 10
        }
    ]
}'
```