package item

import "github.com/diego-dm-morais/order-manager/entity/produto"

type IItem interface {
	EValido() (bool, error)
	GetQuantidade() uint32
	GetProduto() produto.IProduto
}
