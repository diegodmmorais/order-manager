package customer

type ICustomerApi interface {
	FindByCustomer(customerID string) (CustomerOutPutMapper, error)
}
