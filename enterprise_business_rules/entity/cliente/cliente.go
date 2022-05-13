package entity

import (
	"fmt"
	"strings"
)

type cliente struct {
	nome                   string
	telefone               string
	documentoIdentificacao string
}

func (c cliente) EValido() (bool, error) {
	switch {
	case c.nome == "" || strings.Trim(c.nome, " ") == "":
		return false, fmt.Errorf("nome do cliente está vazio")
	case c.telefone == "" || strings.Trim(c.telefone, " ") == "" || len(c.telefone) > 22:
		return false, fmt.Errorf("telefone do cliente está vazio")
	case c.documentoIdentificacao == "" || strings.Trim(c.documentoIdentificacao, " ") == "":
		return false, fmt.Errorf("documento de identificação do cliente está vazio")
	default:
		return true, nil
	}
}
func (c cliente) GetNome() string {
	return c.nome
}
func (c cliente) GetDocumentoIdentificacao() string {
	return c.documentoIdentificacao
}
func (c cliente) GetTelefone() string {
	return c.telefone
}
