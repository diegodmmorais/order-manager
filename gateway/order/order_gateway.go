package order

import (
	"github.com/diego-dm-morais/order-manager/usecase/item"
	"github.com/diego-dm-morais/order-manager/usecase/order"
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
			var itensMapper []ItemInputMapper = make([]ItemInputMapper, len(itens))
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
