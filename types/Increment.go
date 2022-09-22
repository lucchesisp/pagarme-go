package types

type Increment struct {
	Value         int32  `json:"value"`
	Cycles        string `json:"cycles"`
	IncrementType string `json:"increment_type"`
}
