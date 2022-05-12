package customer

type ICustomerGateway interface {
	FindByCustomer(customerID string) (CustomerOutputData, error)
}
