package types

type SplitRule struct {
	amount      string       `json:"amount"`
	recipientId string       `json:"recipientId"`
	ruleType    string       `json:"type"`
	options     SplitOptions `json:"options"`
}
