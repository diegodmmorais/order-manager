package framework

import (
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
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	app._CreateOrderResource(e)
	e.Logger.Fatal(e.Start(HTTP_PORT))
}

func (app Application) _CreateOrderResource(e *echo.Echo) {
	orderRest := framework.CreateOrderRest()
	e.POST("/orders", orderRest.Save)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
