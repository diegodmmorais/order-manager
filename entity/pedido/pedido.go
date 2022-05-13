package entity

import (
	"fmt"

	cliente "github.com/diego-dm-morais/order-manager/entity/cliente"
	endereco "github.com/diego-dm-morais/order-manager/entity/endereco"
	item "github.com/diego-dm-morais/order-manager/entity/item"
)

type pedido struct {
	cliente  cliente.ICliente
	itens    []item.IItem
	endereco endereco.IEndereco
	frete    float32
}

func (p pedido) EValido() (bool, error) {
	switch {
	case p.cliente == nil:
		return false, fmt.Errorf("cliente não foi encontrado")
	case p.itens == nil || len(p.itens) == 0:
		return false, fmt.Errorf("itens não foram encontrados")
	case p.endereco == nil:
		return false, fmt.Errorf("endereço não foi encontrados")
	default:
		clienteValido, clienteError := p.cliente.EValido()
		if !clienteValido {
			return false, clienteError
		}

		enderecoValido, enderecoError := p.endereco.EValido()
		if !enderecoValido {
			return false, enderecoError
		}

		for _, it := range p.itens {
			itemValido, itemError := it.EValido()
			if !itemValido {
				return false, itemError
			}
		}

		return true, nil
	}

}

func (p pedido) GetCliente() cliente.ICliente {
	return p.cliente
}

func (p pedido) GetEnderecoEntrega() endereco.IEndereco {
	return p.endereco
}

func (p pedido) GetItens() []item.IItem {
	return p.itens
}

func (p pedido) GetFrete() float32 {
	return p.frete
}

func (p pedido) GetTotal() float32 {
	var total float32 = 0.0
	for _, it := range p.itens {
		total = total + (it.GetProduto().GetPreco() * float32(it.GetQuantidade()))
	}
	return total + p.frete
}
