package usecase

type ItemRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Amount    uint32 `json:"amount" validate:"required"`
}
