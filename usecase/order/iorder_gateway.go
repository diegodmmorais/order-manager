package order

type IOrderGateway interface {
	save(order OrderDataRequest) (string, error)
}
