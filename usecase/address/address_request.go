package usecase

type ShippingAddressRequest struct {
	Street  string `json:"street"`
	Number  string `json:"number"`
	Zipcode string `json:"zip_code"`
	City    string `json:"city"`
}
