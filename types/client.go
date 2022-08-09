package types

import "time"

// Client is a struct that holds the client information.
type Client struct {
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Code           string    `json:"code"`
	DocumentNumber string    `json:"document_number"`
	DocumentType   string    `json:"document_type"`
	ClientType     string    `json:"client_type"`
	Gender         string    `json:"gender"`
	Address        Address   `json:"address"`
	Phones         Phone     `json:"phones"`
	Birthday       time.Time `json:"birthday"`
	Metadata       string    `json:"metadata"`
}
