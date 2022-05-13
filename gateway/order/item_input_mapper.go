package gateway

type ItemInputMapper struct {
	ProductName string  `bson:"product_name,omitempty"`
	Price       float64 `bson:"price,omitempty"`
	Amount      uint32  `bson:"amount,omitempty"`
}
