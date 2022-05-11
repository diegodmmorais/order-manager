package produto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_produto_e_valido(t *testing.T) {
	t.Parallel()
	var produto IProduto = New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()

	valido, _ := produto.EValido()

	assert.NotNil(t, produto)
	assert.True(t, valido)
	assert.NotNil(t, produto.GetNome())
	assert.NotEmpty(t, produto.GetNome())
	assert.Equal(t, "MacBook Pro 15 2022", produto.GetNome())
	assert.NotNil(t, produto.GetPreco())
	assert.Equal(t, float32(17500.0), produto.GetPreco())
	assert.True(t, produto.EstoqueEstaDisponivel())
}

func Test_produto_com_nome_vazio(t *testing.T) {
	t.Parallel()
	var produto IProduto = New().SetNome("").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()

	valido, mensage := produto.EValido()

	assert.NotNil(t, produto)
	assert.NotNil(t, produto.GetNome())
	assert.Empty(t, produto.GetNome())
	assert.NotNil(t, produto.GetPreco())
	assert.Equal(t, float32(17500.0), produto.GetPreco())
	assert.True(t, produto.EstoqueEstaDisponivel())
	assert.False(t, valido)
	assert.NotNil(t, mensage)
	assert.NotEmpty(t, mensage)
	assert.Equal(t, "nome do produto está vazio", mensage.Error())
}

func Test_produto_com_nome_em_branco(t *testing.T) {
	t.Parallel()
	var produto IProduto = New().SetNome("   ").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()

	valido, mensage := produto.EValido()

	assert.NotNil(t, produto)
	assert.NotNil(t, produto.GetNome())
	assert.NotEmpty(t, produto.GetNome())
	assert.NotNil(t, produto.GetPreco())
	assert.Equal(t, float32(17500.0), produto.GetPreco())
	assert.True(t, produto.EstoqueEstaDisponivel())
	assert.False(t, valido)
	assert.NotNil(t, mensage)
	assert.NotEmpty(t, mensage)
	assert.Equal(t, "nome do produto está vazio", mensage.Error())
}

func Test_produto_com_preco_zero(t *testing.T) {
	t.Parallel()
	var produto IProduto = New().SetNome("MacBook Pro 15 2022").SetEstoqueEstaDisponivel(true).Build()

	valido, mensage := produto.EValido()

	assert.NotNil(t, produto)
	assert.NotNil(t, produto.GetNome())
	assert.NotEmpty(t, produto.GetNome())
	assert.NotNil(t, produto.GetPreco())
	assert.Equal(t, float32(0), produto.GetPreco())
	assert.True(t, produto.EstoqueEstaDisponivel())
	assert.False(t, valido)
	assert.NotNil(t, mensage)
	assert.NotEmpty(t, mensage)
	assert.Equal(t, "preço do produto MacBook Pro 15 2022 está menor ou igual a zero", mensage.Error())
}

func Test_produto_com_estoque_indisponivel(t *testing.T) {
	t.Parallel()
	var produto IProduto = New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(false).Build()

	valido, mensage := produto.EValido()

	assert.NotNil(t, produto)
	assert.NotNil(t, produto.GetNome())
	assert.NotEmpty(t, produto.GetNome())
	assert.NotNil(t, produto.GetPreco())
	assert.Equal(t, float32(17500.0), produto.GetPreco())
	assert.False(t, produto.EstoqueEstaDisponivel())
	assert.False(t, valido)
	assert.NotNil(t, mensage)
	assert.NotEmpty(t, mensage)
	assert.Equal(t, "estoque do produto MacBook Pro 15 2022 não está disponível", mensage.Error())
}
