package types

type SplitOptions struct {
	ChargeProcessingFee bool `json:"chargeProcessingFee"`
	ChargerRemainderFee bool `json:"chargeRemainderFee"`
	Liable              bool `json:"liable"`
}
