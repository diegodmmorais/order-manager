package gateway

import (
	item "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/item"
	order "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"
)

type orderGateway struct {
	order.IOrderGateway
	orderRepository IOrderRepository
}

func (o orderGateway) Save(orderData order.OrderInputData) (string, error) {
	id, error := o.orderRepository.Save(OrderInputMapper{
		IdentificationDocument: orderData.IdentificationDocument,
		Freight:                orderData.Freight,
		Total:                  orderData.Total,
		ShippingAddress: AddressInputMapper{
			Street:  orderData.ShippingAddress.Street,
			Number:  orderData.ShippingAddress.Number,
			Zipcode: orderData.ShippingAddress.Zipcode,
			City:    orderData.ShippingAddress.City,
		},
		Itens: func(itens []item.ItemInputData) []ItemInputMapper {
			var itensMapper []ItemInputMapper
			for _, it := range itens {
				itensMapper = append(itensMapper, ItemInputMapper{
					ProductName: it.ProductName,
					Price:       it.Price,
					Amount:      it.Amount,
				})
			}
			return itensMapper
		}(orderData.Itens),
	})
	return id, error
}
