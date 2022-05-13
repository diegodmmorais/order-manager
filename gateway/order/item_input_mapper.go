package gateway

type ItemInputMapper struct {
	ProductName string  `bson:"product_name"`
	Price       float32 `bson:"price"`
	Amount      uint32  `bson:"amount"`
}
