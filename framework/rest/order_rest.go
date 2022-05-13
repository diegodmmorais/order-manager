package rest

import (
	"net/http"

	orderController "github.com/diego-dm-morais/order-manager/controller/order"
	"github.com/diego-dm-morais/order-manager/usecase/order"
	"github.com/labstack/echo"
)

type OrderRest struct {
	OrderController orderController.IOrderController
}

func (o *OrderRest) Save(c echo.Context) (err error) {
	orderRequest := new(order.OrderRequest)
	if err = c.Bind(orderRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(orderRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	response, erroResponse := o.OrderController.Save(*orderRequest)

	if erroResponse != nil {
		return echo.NewHTTPError(http.StatusBadRequest, erroResponse.Error())
	}

	return c.JSON(http.StatusOK, response)
}
