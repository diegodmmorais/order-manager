package usecase

type IOrderUseCase interface {
	Save(order OrderRequest) (*OrderResponse, error)
}
