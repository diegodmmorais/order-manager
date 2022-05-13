package usecase

import (
	address "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/address"
	item "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/item"
)

type OrderRequest struct {
	CustomerID      string                         `json:"customer_id" validate:"required"`
	ShippingAddress address.ShippingAddressRequest `json:"shipping_address"`
	Freight         float64                        `json:"freight"`
	Items           []item.ItemRequest             `json:"items"`
}
