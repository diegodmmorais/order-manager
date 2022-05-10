package pedido

type IPedido interface {
	GetCliente() ICliente
	GetEnderecoEntrega() IEndereco
	GetItens() []IItens
	GetTotal() float32
}
