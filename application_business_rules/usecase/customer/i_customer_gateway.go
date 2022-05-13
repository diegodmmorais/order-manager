package usecase

type ICustomerGateway interface {
	FindByCustomer(customerID string) (CustomerOutputData, error)
}
