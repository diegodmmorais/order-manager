package produto

type produtoBuilder struct {
	nome                  string
	preco                 float32
	estoqueEstaDisponivel bool
}

func Builder() *produtoBuilder {
	return &produtoBuilder{}
}

func (p *produtoBuilder) SetNome(nome string) *produtoBuilder {
	p.nome = nome
	return p
}

func (p *produtoBuilder) SetPreco(preco float32) *produtoBuilder {
	p.preco = preco
	return p
}

func (p *produtoBuilder) SetEstoqueEstaDisponivel(estoqueEstaDisponivel bool) *produtoBuilder {
	p.estoqueEstaDisponivel = estoqueEstaDisponivel
	return p
}

func (p *produtoBuilder) Build() IProduto {
	return produto{nome: p.nome, preco: p.preco, estoqueEstaDisponivel: p.estoqueEstaDisponivel}
}
