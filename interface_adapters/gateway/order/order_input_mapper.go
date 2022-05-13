package gateway

type OrderInputMapper struct {
	IdentificationDocument string             `bson:"identification_document,omitEmpty"`
	Freight                float64            `bson:"freight,omitempty"`
	Itens                  []ItemInputMapper  `bson:"itens,omitempty"`
	Total                  float64            `bson:"total,omitempty"`
	ShippingAddress        AddressInputMapper `bson:"shipping_address,omitempty"`
}
