package pedido

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/diego-dm-morais/order-manager/entities/produto"
)

func TestSalvarPedido(t *testing.T) {
	t.Parallel()


	var produto produto.IProduto = produto.Builder().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var item IItem = item.Builder().SetProduto(produto).SetQuantidade(1).Build()
	var itens [] IItem = item.Builder().add(item).buildList()
	var endereco IEndereco = endereco.Builder().SetRua("Rua new street").SetNumero("490").SetCep("490098398").SetCidade("São Paulo").Build()
	var cliente ICliente = cliente.Builder().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()
	var pedido IPedido = Builder().SetItens(itens).SetCliente(cliente).SetEnderecoEntrega(endereco).SetFrete(55.9).Build()

	assert.NotNil(t, pedido)
	assert.NotNil(t, pedido.GetItens())
	assert.NotNil(t, pedido.GetCliente())
	assert.NotNil(t, pedido.GetEnderecoEntrega())
	assert.Equal(t, 17555.9, pedido.GetTotal())

	// Validando o Item
	assert.Empty(t, pedido.GetItens())
	assert.NotNil(t, pedido.GetItens()[0])
	assert.Equal(t, true, pedido.GetItens()[0].EValido())
	assert.NotNil(t, pedido.GetItens()[0].GetQuantidade())
	assert.Equal(t, 1, pedido.GetItens()[0].GetQuantidade())

	// Validado o Produto
	assert.NotNil(t, pedido.GetItens()[0].GetProduto())
	assert.Equal(t, true, pedido.GetItens()[0].GetProduto().EValido())
	assert.NotNil(t, pedido.GetItens()[0].GetProduto().GetNome())
	assert.Equal(t, "MacBook Pro 15 2022", pedido.GetItens()[0].GetProduto().GetNome())
	assert.NotNil(t, pedido.GetItens()[0].GetProduto().GetPreco())
	assert.Equal(t, 17500.0, pedido.GetItens()[0].GetProduto().GetPreco())
	assert.True(t, pedido.GetItens()[0].GetProduto().EstoqueEstaDisponivel())

	// Validado o Cliente
	assert.NotNil(t, pedido.GetCliente())
	assert.Equal(t, true, pedido.GetCliente().EValido())
	assert.NotNil(t, pedido.GetCliente().GetNome())
	assert.Equal(t, "Diego Morais", pedido.GetCliente().GetNome())
	assert.NotNil(t, pedido.GetCliente().GetDocumentoIdentificacao())
	assert.Equal(t, "99999999999", pedido.GetCliente().GetDocumentoIdentificacao())
	assert.NotNil(t, pedido.GetCliente().GetTelefone())
	assert.Equal(t, "19 9 98767584", pedido.GetCliente().GetTelefone())

	// Validado a Entrega
	assert.NotNil(t, pedido.GetEnderecoEntrega())
	assert.Equal(t, true, pedido.GetEnderecoEntrega().EValido())
	assert.NotNil(t, pedido.GetEnderecoEntrega().GetRua())
	assert.Equal(t, "Rua new street", pedido.GetEnderecoEntrega().GetRua())
	assert.NotNil(t, pedido.GetEnderecoEntrega().GetNumero())
	assert.Equal(t, "490", pedido.GetEnderecoEntrega().GetNumero())
	assert.NotNil(t, pedido.GetEnderecoEntrega().GetCep())
	assert.Equal(t, "490098398", pedido.GetEnderecoEntrega().GetCep())
	assert.NotNil(t, pedido.GetEnderecoEntrega().GetCidade())
	assert.Equal(t, "São Paulo", pedido.GetEnderecoEntrega().GetCidade())

}
