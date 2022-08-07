package types

import "time"

type Client struct {
	Name           string
	Email          string
	Code           string
	DocumentNumber string
	DocumentType   string
	ClientType     string
	Gender         string
	Address        Address
	Phones         []Phone
	Birthday       time.Time
	Metadata       string
}
