package entity

type enderecoBuilder struct {
	rua    string
	numero string
	cep    string
	cidade string
}

func New() *enderecoBuilder {
	return new(enderecoBuilder)
}

func (e *enderecoBuilder) SetRua(rua string) *enderecoBuilder {
	e.rua = rua
	return e
}

func (e *enderecoBuilder) SetNumero(numero string) *enderecoBuilder {
	e.numero = numero
	return e
}

func (e *enderecoBuilder) SetCep(cep string) *enderecoBuilder {
	e.cep = cep
	return e
}

func (e *enderecoBuilder) SetCidade(cidade string) *enderecoBuilder {
	e.cidade = cidade
	return e
}

func (e *enderecoBuilder) Build() IEndereco {
	return endereco{rua: e.rua, numero: e.numero, cep: e.cep, cidade: e.cidade}
}
