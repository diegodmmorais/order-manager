package order

type IOrderUseCase interface {
	Salvar(order OrderRequest) (*OrderResponse, error)
}
