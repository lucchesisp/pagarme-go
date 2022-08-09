package types

// Phone is a struct that holds the phone number of a customer.
type Phone struct {
	HomePhone   PhoneFormat `json:"home_phone"`
	MobilePhone PhoneFormat `json:"mobile_phone"`
}
