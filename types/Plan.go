package types

// Plan is a struct that holds the plan of a customer.
type Plan struct {
	Name                string            `json:"name"`
	Quantity            int32             `json:"quantity"`
	Description         string            `json:"description"`
	Shippable           bool              `json:"shippable"`
	PaymentMethods      []string          `json:"payment_methods"`
	Installments        []int32           `json:"installments"`
	MinimumPrice        uint32            `json:"minimum_price"`
	StatementDescriptor string            `json:"statement_descriptor"`
	Currency            string            `json:"currency"`
	Interval            string            `json:"interval"`
	IntervalCount       uint8             `json:"interval_count"`
	TrialPeriodDays     uint16            `json:"trial_period_days"`
	BillingType         string            `json:"billing_type"`
	BillingDays         []uint16          `json:"billing_days"`
	PricingScheme       PricingScheme     `json:"pricing_scheme"`
	MetaData            map[string]string `json:"metadata"`
	// TODO: add items field
}
