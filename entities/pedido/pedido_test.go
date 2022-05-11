package pedido

import (
	"testing"

	"github.com/diego-dm-morais/order-manager/entities/cliente"
	"github.com/diego-dm-morais/order-manager/entities/endereco"
	"github.com/diego-dm-morais/order-manager/entities/item"
	"github.com/diego-dm-morais/order-manager/entities/produto"
	"github.com/stretchr/testify/assert"
)

func TestPedidoEValido(t *testing.T) {
	t.Parallel()

	var produto1 produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var produto2 produto.IProduto = produto.New().SetNome("MacBook Pro 13 2022").SetPreco(12300.0).SetEstoqueEstaDisponivel(true).Build()
	var item1 item.IItem = item.New().SetProduto(produto1).SetQuantidade(1).Build()
	var item2 item.IItem = item.New().SetProduto(produto2).SetQuantidade(1).Build()
	var itens []item.IItem = item.New().Add(item1).Add(item2).BuildList()
	var endereco endereco.IEndereco = endereco.New().SetRua("Rua new street").SetNumero("490").SetCep("490098398").SetCidade("São Paulo").Build()
	var cliente cliente.ICliente = cliente.New().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()

	var pedido IPedido = New().SetItens(itens).SetCliente(cliente).SetEnderecoEntrega(endereco).SetFrete(55.9).Build()

	valido, error := pedido.EValido()

	assert.NotNil(t, pedido)
	assert.True(t, valido)
	assert.Nil(t, error)
	assert.Equal(t, float32(55.9), pedido.GetFrete())
	assert.Equal(t, float32(29855.9), pedido.GetTotal())
	assert.NotNil(t, pedido.GetCliente())
	assert.Equal(t, "Diego Morais", pedido.GetCliente().GetNome())
	assert.NotNil(t, pedido.GetEnderecoEntrega())
	assert.Equal(t, "Rua new street", pedido.GetEnderecoEntrega().GetRua())
	assert.NotNil(t, pedido.GetItens())
	assert.NotEmpty(t, pedido.GetItens())
	assert.Equal(t, "MacBook Pro 15 2022", pedido.GetItens()[0].GetProduto().GetNome())

}

func TestPedidoComClienteInvalidoSemDocumentoIdentificacao(t *testing.T) {
	t.Parallel()

	var produto1 produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var produto2 produto.IProduto = produto.New().SetNome("MacBook Pro 13 2022").SetPreco(12300.0).SetEstoqueEstaDisponivel(true).Build()
	var item1 item.IItem = item.New().SetProduto(produto1).SetQuantidade(1).Build()
	var item2 item.IItem = item.New().SetProduto(produto2).SetQuantidade(1).Build()
	var itens []item.IItem = item.New().Add(item1).Add(item2).BuildList()
	var endereco endereco.IEndereco = endereco.New().SetRua("Rua new street").SetNumero("490").SetCep("490098398").SetCidade("São Paulo").Build()
	var cliente cliente.ICliente = cliente.New().SetNome("Diego Morais").SetTelefone("19 9 98767584").Build()

	var pedido IPedido = New().SetItens(itens).SetCliente(cliente).SetEnderecoEntrega(endereco).SetFrete(55.9).Build()

	valido, error := pedido.EValido()

	assert.NotNil(t, pedido)
	assert.False(t, valido)
	assert.Equal(t, "documento de identificação do cliente está vazio", error.Error())

}

func TestPedidoSemCliente(t *testing.T) {
	t.Parallel()

	var produto1 produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var produto2 produto.IProduto = produto.New().SetNome("MacBook Pro 13 2022").SetPreco(12300.0).SetEstoqueEstaDisponivel(true).Build()
	var item1 item.IItem = item.New().SetProduto(produto1).SetQuantidade(1).Build()
	var item2 item.IItem = item.New().SetProduto(produto2).SetQuantidade(1).Build()
	var itens []item.IItem = item.New().Add(item1).Add(item2).BuildList()
	var endereco endereco.IEndereco = endereco.New().SetRua("Rua new street").SetNumero("490").SetCep("490098398").SetCidade("São Paulo").Build()

	var pedido IPedido = New().SetItens(itens).SetEnderecoEntrega(endereco).SetFrete(55.9).Build()

	valido, error := pedido.EValido()

	assert.NotNil(t, pedido)
	assert.False(t, valido)
	assert.Equal(t, "cliente não foi encontrado", error.Error())

}

func TestPedidoComEnderecoSemCep(t *testing.T) {
	t.Parallel()

	var produto1 produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var produto2 produto.IProduto = produto.New().SetNome("MacBook Pro 13 2022").SetPreco(12300.0).SetEstoqueEstaDisponivel(true).Build()
	var item1 item.IItem = item.New().SetProduto(produto1).SetQuantidade(1).Build()
	var item2 item.IItem = item.New().SetProduto(produto2).SetQuantidade(1).Build()
	var itens []item.IItem = item.New().Add(item1).Add(item2).BuildList()
	var endereco endereco.IEndereco = endereco.New().SetRua("Rua new street").SetNumero("490").SetCidade("São Paulo").Build()
	var cliente cliente.ICliente = cliente.New().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()

	var pedido IPedido = New().SetItens(itens).SetCliente(cliente).SetEnderecoEntrega(endereco).SetFrete(55.9).Build()

	valido, error := pedido.EValido()

	assert.NotNil(t, pedido)
	assert.False(t, valido)
	assert.Equal(t, "cep do endereço não foi informada", error.Error())

}

func TestPedidoSemEndereco(t *testing.T) {
	t.Parallel()

	var produto1 produto.IProduto = produto.New().SetNome("MacBook Pro 15 2022").SetPreco(17500.0).SetEstoqueEstaDisponivel(true).Build()
	var produto2 produto.IProduto = produto.New().SetNome("MacBook Pro 13 2022").SetPreco(12300.0).SetEstoqueEstaDisponivel(true).Build()
	var item1 item.IItem = item.New().SetProduto(produto1).SetQuantidade(1).Build()
	var item2 item.IItem = item.New().SetProduto(produto2).SetQuantidade(1).Build()
	var itens []item.IItem = item.New().Add(item1).Add(item2).BuildList()
	var cliente cliente.ICliente = cliente.New().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()

	var pedido IPedido = New().SetItens(itens).SetCliente(cliente).SetFrete(55.9).Build()

	valido, error := pedido.EValido()

	assert.NotNil(t, pedido)
	assert.False(t, valido)
	assert.Equal(t, "endereço não foi encontrados", error.Error())

}

func TestPedidoComItemInvalidoSemProduto(t *testing.T) {
	t.Parallel()

	var produto2 produto.IProduto = produto.New().SetNome("MacBook Pro 13 2022").SetPreco(12300.0).SetEstoqueEstaDisponivel(true).Build()
	var item1 item.IItem = item.New().SetQuantidade(1).Build()
	var item2 item.IItem = item.New().SetProduto(produto2).SetQuantidade(1).Build()
	var itens []item.IItem = item.New().Add(item1).Add(item2).BuildList()
	var endereco endereco.IEndereco = endereco.New().SetRua("Rua new street").SetNumero("490").SetCep("490098398").SetCidade("São Paulo").Build()
	var cliente cliente.ICliente = cliente.New().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()

	var pedido IPedido = New().SetItens(itens).SetCliente(cliente).SetEnderecoEntrega(endereco).SetFrete(55.9).Build()

	valido, error := pedido.EValido()

	assert.NotNil(t, pedido)
	assert.False(t, valido)
	assert.Equal(t, "produto não encontrado", error.Error())

}

func TestPedidoSemItens(t *testing.T) {
	t.Parallel()

	var endereco endereco.IEndereco = endereco.New().SetRua("Rua new street").SetNumero("490").SetCep("490098398").SetCidade("São Paulo").Build()
	var cliente cliente.ICliente = cliente.New().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()

	var pedido IPedido = New().SetCliente(cliente).SetEnderecoEntrega(endereco).SetFrete(55.9).Build()

	valido, error := pedido.EValido()

	assert.NotNil(t, pedido)
	assert.False(t, valido)
	assert.Equal(t, "itens não foram encontrados", error.Error())

}
