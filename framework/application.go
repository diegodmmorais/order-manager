package framework

import (
	"net/http"

	customerApi "github.com/diego-dm-morais/order-manager/apiclient/customer"
	productApi "github.com/diego-dm-morais/order-manager/apiclient/product"
	orderController "github.com/diego-dm-morais/order-manager/controller/order"
	"github.com/diego-dm-morais/order-manager/framework/database"
	"github.com/diego-dm-morais/order-manager/framework/rest"
	"github.com/diego-dm-morais/order-manager/gateway/customer"
	orderGateway "github.com/diego-dm-morais/order-manager/gateway/order"
	"github.com/diego-dm-morais/order-manager/gateway/product"
	"github.com/diego-dm-morais/order-manager/presenter"
	orderRepository "github.com/diego-dm-morais/order-manager/repository/order"
	orderUseCase "github.com/diego-dm-morais/order-manager/usecase/order"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type Application struct {
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

const HTTP_PORT string = ":1323"

func (app *Application) Init() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	app._CreateOrderResource(e)
	e.Logger.Fatal(e.Start(HTTP_PORT))

}

func (app Application) _CreateOrderResource(e *echo.Echo) {

	connectorMongoDataSource := database.CreateConnectorMongoDataSource()
	orderRepository := orderRepository.CreateOrderRepository(connectorMongoDataSource)
	orderGateway := orderGateway.CreateOrderGateway(orderRepository)

	customerApi := customerApi.CreateCustomerApi()
	customerGateway := customer.CreateCustomerGateway(customerApi)
	productApi := productApi.CreateProductApi()
	productGateway := product.CreateProductGateway(productApi)
	orderPresenter := presenter.CreateOrderPresenter()

	orderUseCase := orderUseCase.CreateOrderUseCase(orderPresenter, productGateway, customerGateway, orderGateway)
	orderController := orderController.CreateOrderController(orderUseCase)

	orderRest := new(rest.OrderRest)
	orderRest.OrderController = orderController

	e.POST("/orders", orderRest.Save)

}
