package entity

import (
	"fmt"

	produto "github.com/diego-dm-morais/order-manager/entity/produto"
)

type item struct {
	produto   produto.IProduto
	quatidade uint32
}

func (i item) EValido() (bool, error) {

	switch {
	case nil == i.produto:
		return false, fmt.Errorf("produto não encontrado")
	case i.quatidade == 0:
		return false, fmt.Errorf("quantidade está igual a zero")
	default:
		valido, erro := i.produto.EValido()
		if !valido {
			return false, erro
		}
		return true, nil
	}
}

func (i item) GetQuantidade() uint32 {
	return i.quatidade
}
func (i item) GetProduto() produto.IProduto {
	return i.produto
}
