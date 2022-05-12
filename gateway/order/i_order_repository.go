package order

type IOrderRepository interface {
	Save(orderMapper OrderInputMapper) (string, error)
}
