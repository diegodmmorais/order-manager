package item

type ItemRequest struct {
	ProductID string `json:"product_id"`
	Amount    uint32 `json:"amount"`
}
