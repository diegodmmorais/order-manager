package framework

import (
	customerApi "github.com/diego-dm-morais/order-manager/framework_drivers/apiclient/customer"
	productApi "github.com/diego-dm-morais/order-manager/framework_drivers/apiclient/product"
	orderUseCase "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"
	orderController "github.com/diego-dm-morais/order-manager/interface_adapters/controller"
	database "github.com/diego-dm-morais/order-manager/framework_drivers/database"
	customerGateway "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/customer"
	orderGateway "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/order"
	productGateway "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/product"
	presenter "github.com/diego-dm-morais/order-manager/interface_adapters/presenter"
	orderRepository "github.com/diego-dm-morais/order-manager/interface_adapters/repository/order"
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
