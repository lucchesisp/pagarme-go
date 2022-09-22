package types

// Discount is a struct that represents a discount
type Discount struct {
	Cycles       string `json:"cycles"`
	Value        string `json:"value"`
	DiscountType string `json:"discount_type"`
}
