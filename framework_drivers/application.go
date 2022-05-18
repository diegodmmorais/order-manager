package framework

import (
	"log"
	"net/http"

	framework "github.com/diego-dm-morais/order-manager/framework_drivers/rest"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Application struct {
}

type CustomValidator struct {
	validator *validator.Validate
}

const HTTP_PORT string = ":1323"

func (app *Application) Init() {
	log.Println("Iniciando a aplicação")
	e := echo.New()

	log.Println("Criando orquestrador de logger")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${protocol} ${method} '${host}${uri}' status=${status}\n",
	}))

	log.Println("Criando o validador de dados json")
	e.Validator = &CustomValidator{validator: validator.New()}

	log.Println("Criando grupo: /api/v1")
	group := e.Group("/api/v1")

	log.Println("Criando o serviço order rest api ")
	app._CreateOrderResource(group)

	e.Logger.Fatal(e.Start(HTTP_PORT))
	log.Println("Aplicação iniciada")
}

func (app Application) _CreateOrderResource(e *echo.Group) {
	orderRest := framework.CreateOrderRest()
	e.POST("/orders", orderRest.Save)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
