package entity

import (
	cliente "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/cliente"
	endereco "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/endereco"
	item "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/item"
)

type pedidoBuilder struct {
	cliente  cliente.ICliente
	itens    []item.IItem
	endereco endereco.IEndereco
	frete    float64
}

func New() *pedidoBuilder {
	return new(pedidoBuilder)
}

func (p *pedidoBuilder) SetCliente(cliente cliente.ICliente) *pedidoBuilder {
	p.cliente = cliente
	return p
}

func (p *pedidoBuilder) SetItens(itens []item.IItem) *pedidoBuilder {
	p.itens = itens
	return p
}

func (p *pedidoBuilder) SetEnderecoEntrega(endereco endereco.IEndereco) *pedidoBuilder {
	p.endereco = endereco
	return p
}

func (p *pedidoBuilder) SetFrete(frete float64) *pedidoBuilder {
	p.frete = frete
	return p
}

func (p *pedidoBuilder) Build() IPedido {
	return pedido{cliente: p.cliente, itens: p.itens, endereco: p.endereco, frete: p.frete}
}
