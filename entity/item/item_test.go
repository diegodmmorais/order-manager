package item

import (
	"testing"

	"github.com/diego-dm-morais/order-manager/entity/produto"
	"github.com/stretchr/testify/assert"
)

func Test_iten_e_valido(t *testing.T) {
	t.Parallel()

	var produto produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var item IItem = New().SetProduto(produto).SetQuantidade(1).Build()
	valido, erro := item.EValido()

	assert.NotNil(t, item)
	assert.NotNil(t, item.GetQuantidade())
	assert.Equal(t, uint32(1), item.GetQuantidade())
	assert.NotNil(t, item.GetProduto())
	assert.True(t, valido)
	assert.Nil(t, erro)
}

func Test_iten_com_produto_invalido(t *testing.T) {
	t.Parallel()

	var produto produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetEstoqueEstaDisponivel(true).Build()
	var item IItem = New().SetProduto(produto).SetQuantidade(1).Build()
	valido, erro := item.EValido()

	assert.NotNil(t, item)
	assert.NotNil(t, item.GetQuantidade())
	assert.Equal(t, uint32(1), item.GetQuantidade())
	assert.NotNil(t, item.GetProduto())
	assert.False(t, valido)
	assert.Equal(t, "preço do produto MacBook Pro 15 2022 está menor ou igual a zero", erro.Error())
}

func Test_iten_sem_produto(t *testing.T) {
	t.Parallel()

	var item IItem = New().SetQuantidade(1).Build()
	valido, erro := item.EValido()

	assert.NotNil(t, item)
	assert.NotNil(t, item.GetQuantidade())
	assert.Equal(t, uint32(1), item.GetQuantidade())
	assert.Nil(t, item.GetProduto())
	assert.False(t, valido)
	assert.Equal(t, "produto não encontrado", erro.Error())
}

func Test_iten_quantidade_disponivel(t *testing.T) {
	t.Parallel()

	var produto produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var item IItem = New().SetProduto(produto).Build()
	valido, erro := item.EValido()

	assert.NotNil(t, item)
	assert.NotNil(t, item.GetQuantidade())
	assert.Equal(t, uint32(0), item.GetQuantidade())
	assert.NotNil(t, item.GetProduto())
	assert.False(t, valido)
	assert.Equal(t, "quantidade está igual a zero", erro.Error())
}

func Test_itens_validos(t *testing.T) {
	t.Parallel()

	var produto produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var item IItem = New().SetProduto(produto).SetQuantidade(1).Build()
	var itens []IItem = New().Add(item).BuildList()

	for _, it := range itens {
		valido, erro := item.EValido()
		assert.NotNil(t, it)
		assert.NotNil(t, it.GetQuantidade())
		assert.Equal(t, uint32(1), it.GetQuantidade())
		assert.NotNil(t, it.GetProduto())
		assert.True(t, valido)
		assert.Nil(t, erro)
	}
}
