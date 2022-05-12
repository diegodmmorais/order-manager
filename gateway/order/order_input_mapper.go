package order

type OrderInputMapper struct {
	IdentificationDocument string
	Freight                float32
	Itens                  []ItemInputMapper
	Total                  float32
	ShippingAddress        AddressInputMapper
}
