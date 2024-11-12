package model

type Order struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Diskon      float64 `json:"diskon"`
	Amount      float64 `json:"amount"`
	FinalAmount float64 `json:"final_amount"`
}
