package pedido

import (
	"github.com/diego-dm-morais/order-manager/entity/cliente"
	"github.com/diego-dm-morais/order-manager/entity/endereco"
	"github.com/diego-dm-morais/order-manager/entity/item"
)

type IPedido interface {
	EValido() (bool, error)
	GetCliente() cliente.ICliente
	GetEnderecoEntrega() endereco.IEndereco
	GetItens() []item.IItem
	GetTotal() float32
	GetFrete() float32
}
