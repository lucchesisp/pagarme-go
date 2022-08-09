package types

// PhoneFormat is a struct that holds the phone number of a customer.
type PhoneFormat struct {
	CountryCode string `json:"country_code"`
	AreaCode    string `json:"area_code"`
	Number      string `json:"number"`
}
