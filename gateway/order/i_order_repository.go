package gateway

type IOrderRepository interface {
	Save(orderMapper OrderInputMapper) (string, error)
}
