package framework

import (
	"log"
	"net/http"

	framework "github.com/diego-dm-morais/order-manager/framework_drivers/rest"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type Application struct {
}

type CustomValidator struct {
	validator *validator.Validate
}

const HTTP_PORT string = ":1323"

func (app *Application) Init() {
	log.Println("Iniciando Iniciando o serviço")
	e := echo.New()
	log.Println("Criando o validador de dados json")
	e.Validator = &CustomValidator{validator: validator.New()}
	app._CreateOrderResource(e)
	e.Logger.Fatal(e.Start(HTTP_PORT))
}

func (app Application) _CreateOrderResource(e *echo.Echo) {
	log.Println("Criando o serviço order rest api ")
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
