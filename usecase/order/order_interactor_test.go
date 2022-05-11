package order

import (
	"fmt"
	"testing"

	"github.com/diego-dm-morais/order-manager/usecase/address"
	"github.com/diego-dm-morais/order-manager/usecase/customer"
	"github.com/diego-dm-morais/order-manager/usecase/item"
	"github.com/diego-dm-morais/order-manager/usecase/product"
)

func TestSalvarPedido(t *testing.T) {
	t.Parallel()

	var orderInputBoundary IOrderInputBoundary = CreateOrderInputBoundary(MockOrderOutputBoundary{}, product.MockProductGateway{}, customer.MockCustomerGateway{}, MockOrderGateway{})

	response, erro := orderInputBoundary.Salvar(OrderRequest{
		CustomerID: "1",
		Freight:    55.6,
		Items: []item.ItemRequest{
			{
				ProductID: "1",
				Amount:    2,
			},
		},
		ShippingAddress: address.ShippingAddressRequest{
			Street:  "Rua teste",
			Number:  "10",
			Zipcode: "750989899",
			City:    "Goiania",
		},
	})

	fmt.Println(response, erro)

}
