package types

// PriceBrackets is a struct that represents a price bracket
type PriceBrackets []struct {
	StartQuantity int32 `json:"start_quantity"`
	EndQuantity   int32 `json:"end_quantity"`
	OveragePrice  int32 `json:"overage_price"`
	Price         int32 `json:"price"`
}
