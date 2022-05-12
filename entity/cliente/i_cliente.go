package cliente

type ICliente interface {
	EValido() (bool, error)
	GetNome() string
	GetDocumentoIdentificacao() string
	GetTelefone() string
}
