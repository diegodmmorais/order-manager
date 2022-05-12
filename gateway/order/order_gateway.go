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
	orderMapperResponse, error := o.orderRepository.Save(OrderMapperRequest{
		IdentificationDocument: orderData.IdentificationDocument,
		Freight:                orderData.Freight,
		Total:                  orderData.Total,
		ShippingAddress: ShippingAddressMapperRequest{
			Street:  orderData.ShippingAddress.Street,
			Number:  orderData.ShippingAddress.Number,
			Zipcode: orderData.ShippingAddress.Zipcode,
			City:    orderData.ShippingAddress.City,
		},
		Itens: func(itens []item.ItemInputData) []ItemMapperRequest {
			var itensMapper []ItemMapperRequest = make([]ItemMapperRequest, len(itens))
			for _, it := range itens {
				itensMapper = append(itensMapper, ItemMapperRequest{
					ProductName: it.ProductName,
					Price:       it.Price,
					Amount:      it.Amount,
				})
			}
			return itensMapper
		}(orderData.Itens),
	})
	return orderMapperResponse.ID, error
}
