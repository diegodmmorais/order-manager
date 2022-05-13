package entity

import (
	cliente "github.com/diego-dm-morais/order-manager/entity/cliente"
	endereco "github.com/diego-dm-morais/order-manager/entity/endereco"
	item "github.com/diego-dm-morais/order-manager/entity/item"
)

type IPedido interface {
	EValido() (bool, error)
	GetCliente() cliente.ICliente
	GetEnderecoEntrega() endereco.IEndereco
	GetItens() []item.IItem
	GetTotal() float64
	GetFrete() float64
}
