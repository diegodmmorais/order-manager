package order

import (
	"github.com/diego-dm-morais/order-manager/usecase/address"
	"github.com/diego-dm-morais/order-manager/usecase/item"
)

type OrderInputData struct {
	IdentificationDocument string
	Freight                float32
	Itens                  []item.ItemInputData
	Total                  float32
	ShippingAddress        address.AddressInputData
}
