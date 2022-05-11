package endereco

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnderecoEValido(t *testing.T) {
	t.Parallel()

	var endereco IEndereco = New().SetRua("Rua new street").SetNumero("490").SetCep("490098398").SetCidade("São Paulo").Build()
	valido, erro := endereco.EValido()

	assert.NotNil(t, endereco)
	assert.True(t, valido)
	assert.Nil(t, erro)
	assert.NotNil(t, endereco.GetCep())
	assert.NotEmpty(t, endereco.GetCep())
	assert.Equal(t, "490098398", endereco.GetCep())

	assert.NotNil(t, endereco.GetRua())
	assert.NotEmpty(t, endereco.GetRua())
	assert.Equal(t, "Rua new street", endereco.GetRua())

	assert.NotNil(t, endereco.GetNumero())
	assert.NotEmpty(t, endereco.GetNumero())
	assert.Equal(t, "490", endereco.GetNumero())

	assert.NotNil(t, endereco.GetCidade())
	assert.NotEmpty(t, endereco.GetCidade())
	assert.Equal(t, "São Paulo", endereco.GetCidade())

}

func TestEnderecoSemRua(t *testing.T) {
	t.Parallel()

	var endereco IEndereco = New().SetNumero("490").SetCep("490098398").SetCidade("São Paulo").Build()
	valido, erro := endereco.EValido()

	assert.NotNil(t, endereco)
	assert.False(t, valido)
	assert.NotNil(t, erro)
	assert.Equal(t, "rua do endereço não foi informada", erro.Error())
}

func TestEnderecoSemNumero(t *testing.T) {
	t.Parallel()

	var endereco IEndereco = New().SetRua("Rua new street").SetCep("490098398").SetCidade("São Paulo").Build()
	valido, erro := endereco.EValido()

	assert.NotNil(t, endereco)
	assert.False(t, valido)
	assert.NotNil(t, erro)
	assert.Equal(t, "numero do endereço não foi informada", erro.Error())
}

func TestEnderecoSemCep(t *testing.T) {
	t.Parallel()

	var endereco IEndereco = New().SetRua("Rua new street").SetNumero("490").SetCidade("São Paulo").Build()
	valido, erro := endereco.EValido()

	assert.NotNil(t, endereco)
	assert.False(t, valido)
	assert.NotNil(t, erro)
	assert.Equal(t, "cep do endereço não foi informada", erro.Error())
}

func TestEnderecoSemCidade(t *testing.T) {
	t.Parallel()

	var endereco IEndereco = New().SetRua("Rua new street").SetNumero("490").SetCep("490098398").Build()
	valido, erro := endereco.EValido()

	assert.NotNil(t, endereco)
	assert.False(t, valido)
	assert.NotNil(t, erro)
	assert.Equal(t, "cidade do endereço não foi informada", erro.Error())
}
