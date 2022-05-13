package entity

type IProduto interface {
	EValido() (bool, error)
	GetNome() string
	GetPreco() float32
	EstoqueEstaDisponivel() bool
}
