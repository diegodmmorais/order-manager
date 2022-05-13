package entity

type IEndereco interface {
	EValido() (bool, error)
	GetRua() string
	GetNumero() string
	GetCep() string
	GetCidade() string
}
