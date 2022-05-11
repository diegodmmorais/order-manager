package item

import "github.com/diego-dm-morais/order-manager/entities/produto"

type itemBuilder struct {
	produto   produto.IProduto
	quatidade uint32
	itens     []IItem
}

func New() *itemBuilder {
	return &itemBuilder{}
}

func (i *itemBuilder) SetProduto(produto produto.IProduto) *itemBuilder {
	i.produto = produto
	return i
}

func (i *itemBuilder) SetQuantidade(quantidade uint32) *itemBuilder {
	i.quatidade = quantidade
	return i
}

func (i *itemBuilder) Add(item IItem) *itemBuilder {
	i.itens = append(i.itens, item)
	return i
}

func (i *itemBuilder) Build() IItem {
	return item{quatidade: i.quatidade, produto: i.produto}
}

func (i *itemBuilder) BuildList() []IItem {
	return i.itens
}
