package cliente

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClienteEValido(t *testing.T) {
	t.Parallel()
	var cliente ICliente = Builder().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()

	valido, error := cliente.EValido()

	assert.NotNil(t, cliente)
	assert.True(t, valido)
	assert.Nil(t, error)

	assert.NotNil(t, cliente.GetNome())
	assert.NotEmpty(t, cliente.GetNome())
	assert.Equal(t, "Diego Morais", cliente.GetNome())

	assert.NotNil(t, cliente.GetTelefone())
	assert.NotEmpty(t, cliente.GetTelefone())
	assert.Equal(t, "19 9 98767584", cliente.GetTelefone())

	assert.NotNil(t, cliente.GetDocumentoIdentificacao())
	assert.NotEmpty(t, cliente.GetDocumentoIdentificacao())
	assert.Equal(t, "99999999999", cliente.GetDocumentoIdentificacao())
}

func TestClienteComNomeInvalido(t *testing.T) {
	t.Parallel()
	var cliente1 ICliente = Builder().SetNome("").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()
	var cliente2 ICliente = Builder().SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()
	var cliente3 ICliente = Builder().SetNome("  ").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("99999999999").Build()

	valido1, error1 := cliente1.EValido()
	valido2, error2 := cliente2.EValido()
	valido3, error3 := cliente2.EValido()

	assert.NotNil(t, cliente1)
	assert.False(t, valido1)
	assert.Equal(t, "nome do cliente está vazio",  error1.Error())

	assert.NotNil(t, cliente2)
	assert.False(t, valido2)
	assert.Equal(t, "nome do cliente está vazio",  error2.Error())

	assert.NotNil(t, cliente3)
	assert.False(t, valido3)
	assert.Equal(t, "nome do cliente está vazio",  error3.Error())

}


func TestClienteComTelefoneInvalido(t *testing.T) {
	t.Parallel()
	var cliente1 ICliente = Builder().SetNome("Diego Morais").SetDocumentoIdentificacao("99999999999").Build()
	var cliente2 ICliente = Builder().SetNome("Diego Morais").SetTelefone("").SetDocumentoIdentificacao("99999999999").Build()
	var cliente3 ICliente = Builder().SetNome("Diego Morais").SetTelefone("  ").SetDocumentoIdentificacao("99999999999").Build()

	valido1, error1 := cliente1.EValido()
	valido2, error2 := cliente2.EValido()
	valido3, error3 := cliente2.EValido()

	assert.NotNil(t, cliente1)
	assert.False(t, valido1)
	assert.Equal(t, "telefone do cliente está vazio",  error1.Error())

	assert.NotNil(t, cliente2)
	assert.False(t, valido2)
	assert.Equal(t, "telefone do cliente está vazio",  error2.Error())

	assert.NotNil(t, cliente3)
	assert.False(t, valido3)
	assert.Equal(t, "telefone do cliente está vazio",  error3.Error())

}

func TestClienteComDocumentoInvalido(t *testing.T) {
	t.Parallel()
	var cliente1 ICliente = Builder().SetNome("Diego Morais").SetTelefone("19 9 98767584").Build()
	var cliente2 ICliente = Builder().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("").Build()
	var cliente3 ICliente = Builder().SetNome("Diego Morais").SetTelefone("19 9 98767584").SetDocumentoIdentificacao("  ").Build()

	valido1, error1 := cliente1.EValido()
	valido2, error2 := cliente2.EValido()
	valido3, error3 := cliente2.EValido()

	assert.NotNil(t, cliente1)
	assert.False(t, valido1)
	assert.Equal(t, "documento de identificação do cliente está vazio",  error1.Error())

	assert.NotNil(t, cliente2)
	assert.False(t, valido2)
	assert.Equal(t, "documento de identificação do cliente está vazio",  error2.Error())

	assert.NotNil(t, cliente3)
	assert.False(t, valido3)
	assert.Equal(t, "documento de identificação do cliente está vazio",  error3.Error())

}

