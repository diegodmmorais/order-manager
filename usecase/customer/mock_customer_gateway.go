package customer

type MockCustomerGateway struct {
	ICustomerGateway
}

func (c MockCustomerGateway) FindByCustomer(customerID string) (CustomerResponse, error) {
	return CustomerResponse{
		Name:                   "Diego",
		Telephone:              "19 9 9988 9989",
		IdentificationDocument: "99999999999",
	}, nil
}
