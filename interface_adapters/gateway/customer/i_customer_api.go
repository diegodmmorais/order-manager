package gateway

type ICustomerApi interface {
	FindByCustomer(customerID string) (CustomerOutPutMapper, error)
}
