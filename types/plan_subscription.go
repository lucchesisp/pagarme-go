package types

// PlanSubscription is a struct that represents a plan subscription
type PlanSubscription struct {
	Code                 string       `json:"code"`
	PlanID               string       `json:"plan_id"`
	PaymentMethod        string       `json:"payment_method"`
	StartAt              string       `json:"start_at"`
	ClientID             string       `json:"customer_id"`
	Client               *Client      `json:"customer"`
	Card                 *Card        `json:"card"`
	Discount             *[]Discount  `json:"discount"`
	Increments           *[]Increment `json:"increments"`
	GatewayAffiliationID string       `json:"gateway_affiliation_id"`
	BoletoDueDays        string       `json:"boleto_due_days"`
}
