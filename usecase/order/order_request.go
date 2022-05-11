package order

import (
	"github.com/diego-dm-morais/order-manager/usecase/address"
	"github.com/diego-dm-morais/order-manager/usecase/item"
)

type OrderRequest struct {
	CustomerID      string                         `json:"customer_id"`
	ShippingAddress address.ShippingAddressRequest `json:"shipping_address"`
	Freight         float32                        `json:"freight"`
	Items           []item.ItemRequest             `json:"items"`
}
