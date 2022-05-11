package customer

type ICustomerGateway interface {
	FindByCustomer(customerID string) CustomerResponse
}
