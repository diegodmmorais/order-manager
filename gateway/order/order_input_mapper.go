package order

type OrderInputMapper struct {
	IdentificationDocument string             `bson:"identification_document"`
	Freight                float32            `bson:"freight"`
	Itens                  []ItemInputMapper  `bson:"itens"`
	Total                  float32            `bson:"total"`
	ShippingAddress        AddressInputMapper `bson:"shipping_address"`
}
