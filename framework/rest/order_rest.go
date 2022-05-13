package framework

import (
	"net/http"

	orderController "github.com/diego-dm-morais/order-manager/controller"
	orderUseCase "github.com/diego-dm-morais/order-manager/usecase/order"
	"github.com/labstack/echo"
)

type orderRest struct {
	orderController orderController.IOrderController
}

func (o *orderRest) Save(c echo.Context) (err error) {
	orderRequest := new(orderUseCase.OrderRequest)
	if err = c.Bind(orderRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(orderRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	response, erroResponse := o.orderController.Save(*orderRequest)

	if erroResponse != nil {
		return echo.NewHTTPError(http.StatusBadRequest, erroResponse.Error())
	}

	return c.JSON(http.StatusOK, response)
}
