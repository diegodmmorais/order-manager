package usecase

import (
	address "github.com/diego-dm-morais/order-manager/usecase/address"
	item "github.com/diego-dm-morais/order-manager/usecase/item"
)

type OrderInputData struct {
	IdentificationDocument string
	Freight                float64
	Itens                  []item.ItemInputData
	Total                  float64
	ShippingAddress        address.AddressInputData
}
