package framework

import (
	customerApi "github.com/diego-dm-morais/order-manager/apiclient/customer"
	productApi "github.com/diego-dm-morais/order-manager/apiclient/product"
	orderController "github.com/diego-dm-morais/order-manager/controller"
	database "github.com/diego-dm-morais/order-manager/framework/database"
	customerGateway "github.com/diego-dm-morais/order-manager/gateway/customer"
	orderGateway "github.com/diego-dm-morais/order-manager/gateway/order"
	productGateway "github.com/diego-dm-morais/order-manager/gateway/product"
	presenter "github.com/diego-dm-morais/order-manager/presenter"
	orderRepository "github.com/diego-dm-morais/order-manager/repository/order"
	orderUseCase "github.com/diego-dm-morais/order-manager/usecase/order"
)

func CreateOrderRest() IOrderRest {
	connectorMongoDataSource := database.CreateConnectorMongoDataSource()
	orderRepository := orderRepository.CreateOrderRepository(connectorMongoDataSource)
	orderGateway := orderGateway.CreateOrderGateway(orderRepository)

	customerApi := customerApi.CreateCustomerApi()
	customerGateway := customerGateway.CreateCustomerGateway(customerApi)
	productApi := productApi.CreateProductApi()
	productGateway := productGateway.CreateProductGateway(productApi)
	orderPresenter := presenter.CreateOrderPresenter()

	orderUseCase := orderUseCase.CreateOrderUseCase(orderPresenter, productGateway, customerGateway, orderGateway)
	orderController := orderController.CreateOrderController(orderUseCase)

	return &orderRest{orderController: orderController}
}
