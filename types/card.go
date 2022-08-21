package types

type Card struct {
	CustomerID      string   `json:"customer_id"`
	Number          string   `json:"number"`
	HolderName      string   `json:"holder_name"`
	HolderDocument  string   `json:"holder_document"`
	ExpirationMonth int32    `json:"exp_month"`
	ExpirationYear  int32    `json:"exp_year"`
	CVV             string   `json:"cvv"`
	Brand           string   `json:"brand"`
	Label           string   `json:"label"`
	BillingAddress  *Address `json:"billing_address"`
}
