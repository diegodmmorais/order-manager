package entity

type produtoBuilder struct {
	nome                  string
	preco                 float64
	estoqueEstaDisponivel bool
}

func New() *produtoBuilder {
	return new(produtoBuilder)
}

func (p *produtoBuilder) SetNome(nome string) *produtoBuilder {
	p.nome = nome
	return p
}

func (p *produtoBuilder) SetPreco(preco float64) *produtoBuilder {
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
