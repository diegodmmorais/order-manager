package usecase

import (
	"log"

	address "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/address"
	customer "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/customer"
	itemUseCase "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/item"
	product "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/product"
	cliente "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/cliente"
	endereco "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/endereco"
	item "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/item"
	pedido "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/pedido"
	produto "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/produto"
)

type orderUseCase struct {
	IOrderUseCase
	orderOutputBoundary IOrderPresenter
	productGateway      product.IProductGateway
	customerGateway     customer.ICustomerGateway
	orderGateway        IOrderGateway
}

func (o orderUseCase) Save(order OrderRequest) (*OrderResponse, error) {

	channel := make(chan interface{}, 3)

	go o._Load(order, channel)
	
	itens := (<-channel).([]item.IItem)
	cliente := (<-channel).(cliente.ICliente)
	endereco := (<-channel).(endereco.IEndereco)

	var pedido pedido.IPedido = pedido.New().SetItens(itens).SetCliente(cliente).SetEnderecoEntrega(endereco).SetFrete(order.Freight).Build()

	_, erro := pedido.EValido()
	if erro != nil {
		log.Println(erro)
		return nil, erro
	}

	var orderInputData OrderInputData = OrderInputData{
		IdentificationDocument: cliente.GetDocumentoIdentificacao(),
		Freight:                pedido.GetFrete(),
		Total:                  pedido.GetTotal(),
		Itens:                  o._GetItensData(itens),
		ShippingAddress: address.AddressInputData{
			Street:  endereco.GetRua(),
			Number:  endereco.GetNumero(),
			Zipcode: endereco.GetCep(),
			City:    endereco.GetCidade(),
		},
	}

	orderID, erroData := o.orderGateway.Save(orderInputData)
	if erroData != nil {
		log.Println(erroData)
		return nil, erroData
	}

	return o.orderOutputBoundary.Success(OrderSuccessInputData{OrderID: orderID, CustomerName: cliente.GetNome()})
}

func (o orderUseCase) _Load(order OrderRequest, channel chan any) {
	channel <- o._GetItens(order.Items)
	channel <- o._GetCustomer(order.CustomerID)
	channel <- o._GetAddress(order.ShippingAddress)
}

func (o orderUseCase) _GetItens(itensRequest []itemUseCase.ItemRequest) []item.IItem {
	log.Println("Buscando os produtos")
	var itens []item.IItem
	for _, it := range itensRequest {
		product, _ := o.productGateway.FindByProduct(it.ProductID)
		var produto produto.IProduto = produto.New().SetNome(product.Nome).SetPreco(product.Price).SetEstoqueEstaDisponivel(product.EstoqueEstaDisponivel).Build()
		var item item.IItem = item.New().SetProduto(produto).SetQuantidade(it.Amount).Build()
		itens = append(itens, item)
	}
	return itens
}

func (o orderUseCase) _GetCustomer(customerID string) cliente.ICliente {
	log.Printf("Buscando o cliente %s", customerID)
	customer, _ := o.customerGateway.FindByCustomer(customerID)
	return cliente.New().SetNome(customer.Name).SetDocumentoIdentificacao(customer.IdentificationDocument).SetTelefone(customer.Telephone).Build()
}

func (o orderUseCase) _GetAddress(address address.ShippingAddressRequest) endereco.IEndereco {
	return endereco.New().SetCep(address.Zipcode).SetCidade(address.City).SetNumero(address.Number).SetRua(address.Street).Build()
}

func (o orderUseCase) _GetItensData(itens []item.IItem) []itemUseCase.ItemInputData {
	var itensData []itemUseCase.ItemInputData
	for _, it := range itens {
		itensData = append(itensData, itemUseCase.ItemInputData{
			ProductName: it.GetProduto().GetNome(),
			Price:       it.GetProduto().GetPreco(),
			Amount:      it.GetQuantidade(),
		})
	}
	return itensData
}
