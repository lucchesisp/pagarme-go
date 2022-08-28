package types

type PricingScheme struct {
	SchemeType    string          `json:"scheme_type"`
	Price         int32           `json:"price"`
	MinimumPrice  int32           `json:"minimum_price"`
	PriceBrackets []PriceBrackets `json:"price_brackets"`
}
