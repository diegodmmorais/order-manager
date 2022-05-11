package cliente

type clienteBuilder struct {
	nome                   string
	telefone               string
	documentoIdentificacao string
}

func New() *clienteBuilder {
	return &clienteBuilder{}
}

func (c *clienteBuilder) SetNome(nome string) *clienteBuilder {
	c.nome = nome
	return c
}

func (c *clienteBuilder) SetDocumentoIdentificacao(documentoIdentificacao string) *clienteBuilder {
	c.documentoIdentificacao = documentoIdentificacao
	return c
}

func (c *clienteBuilder) SetTelefone(telefone string) *clienteBuilder {
	c.telefone = telefone
	return c
}

func (c *clienteBuilder) Build() ICliente {
	return cliente{nome: c.nome, telefone: c.telefone, documentoIdentificacao: c.documentoIdentificacao}
}
