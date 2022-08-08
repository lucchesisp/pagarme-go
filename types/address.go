package types

type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	Line1   string `json:"line_1"`
	ZipCode string `json:"zip_code"`
}
