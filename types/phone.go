package types

type Phone struct {
	HomePhone   PhoneFormat `json:"home_phone"`
	MobilePhone PhoneFormat `json:"mobile_phone"`
}
