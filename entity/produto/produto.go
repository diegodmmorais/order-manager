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
	case p.nome == "" || strings.Trim(p.nome, " ") == "":
		return false, fmt.Errorf("nome do produto está vazio")
	case p.preco <= 0:
		return false, fmt.Errorf("preço do produto %s está menor ou igual a zero", p.nome)
	case !p.estoqueEstaDisponivel:
		return false, fmt.Errorf("estoque do produto %s não está disponível", p.nome)
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
