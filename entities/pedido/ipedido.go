package pedido

import (
	"github.com/diego-dm-morais/order-manager/entities/cliente"
	"github.com/diego-dm-morais/order-manager/entities/endereco"
	"github.com/diego-dm-morais/order-manager/entities/item"
)

type IPedido interface {
	EValido() (bool, error)
	GetCliente() cliente.ICliente
	GetEnderecoEntrega() endereco.IEndereco
	GetItens() []item.IItem
	GetTotal() float32
	GetFrete() float32
}
