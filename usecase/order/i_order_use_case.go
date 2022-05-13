package order

type IOrderUseCase interface {
	Save(order OrderRequest) (*OrderResponse, error)
}
