package usecase

type OrderResponse struct {
	ID           string `json:"id"`
	Message      string `json:"message"`
	CustomerName string `json:"customer_name"`
}
