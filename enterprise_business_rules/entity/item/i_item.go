package entity

import produto "github.com/diego-dm-morais/order-manager/enterprise_business_rules/entity/produto"

type IItem interface {
	EValido() (bool, error)
	GetQuantidade() uint32
	GetProduto() produto.IProduto
}
