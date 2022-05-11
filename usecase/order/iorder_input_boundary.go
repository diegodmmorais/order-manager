package order

type IOrderInputBoundary interface {
	Salvar(order OrderRequest) (*OrderResponse, error)
}
