package customer

type mockCustomerApi struct {
	ICustomerApi
}

func (c mockCustomerApi) FindByCustomer(customerID string) (CustomerResponseMapper, error) {
	return CustomerResponseMapper{
		Name:                   "Diego Morais",
		Telephone:              "+55 015 11 9 9999-9999",
		IdentificationDocument: "000.000.000-00",
	}, nil
}
