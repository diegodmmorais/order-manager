package order

type IOrderRepository interface {
	Save(orderMapper OrderMapperRequest) (OrderMapperResponse, error)
}