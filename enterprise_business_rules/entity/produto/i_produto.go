package entity

type IProduto interface {
	EValido() (bool, error)
	GetNome() string
	GetPreco() float64
	EstoqueEstaDisponivel() bool
}
