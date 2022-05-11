package endereco

import (
	"fmt"
	"strings"
)

type endereco struct {
	rua    string
	numero string
	cep    string
	cidade string
}

func (e endereco) EValido() (bool, error) {
	switch {
	case e.rua == "" || strings.Trim(e.rua, " ") == "":
		return false, fmt.Errorf("rua do endereço não foi informada")
	case e.numero == "" || strings.Trim(e.numero, " ") == "":
		return false, fmt.Errorf("numero do endereço não foi informada")
	case e.cep == "" || strings.Trim(e.cep, " ") == "":
		return false, fmt.Errorf("cep do endereço não foi informada")
	case e.cidade == "" || strings.Trim(e.cidade, " ") == "":
		return false, fmt.Errorf("cidade do endereço não foi informada")
	default:
		return true, nil
	}
}
func (e endereco) GetRua() string {
	return e.rua
}
func (e endereco) GetNumero() string {
	return e.numero
}
func (e endereco) GetCep() string {
	return e.cep
}
func (e endereco) GetCidade() string {
	return e.cidade
}
