package customer

type ICustomerApi interface{
	FindByCustomer(customerID string) (CustomerResponseMapper, error)
}