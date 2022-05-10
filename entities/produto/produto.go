package produto

import (
	"fmt"
	"strings"
)

type produto struct {
	nome                  string
	preco                 float32
	estoqueEstaDisponivel bool
}

func (p produto) EValido() (bool, error) {
	switch {
	case "" == p.nome || "" == strings.Trim(p.nome, " "):
		return false, fmt.Errorf("Nome do produto está vazio")
	case 0 >= p.preco:
		return false, fmt.Errorf("Preço do produto %s está menor ou igual a zero", p.nome)
	case !p.estoqueEstaDisponivel:
		return false, fmt.Errorf("Estoque do produto %s não está disponível", p.nome)
	default:
		return true, nil
	}
}
func (p produto) GetNome() string {
	return p.nome
}
func (p produto) GetPreco() float32 {
	return p.preco
}

func (p produto) EstoqueEstaDisponivel() bool {
	return p.estoqueEstaDisponivel
}
