package order

type OrderMapperRequest struct {
	IdentificationDocument string
	Freight                float32
	Itens                  []ItemMapperRequest
	Total                  float32
	ShippingAddress        ShippingAddressMapperRequest
}
