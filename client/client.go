package client

import (
	"encoding/json"
	"github.com/lucchesisp/pagarme-go/enums"
	"github.com/lucchesisp/pagarme-go/utils"
	"time"
)

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

type Address struct {
	Country string
	State   string
	City    string
	ZipCode string
}

type Phone struct {
	HomePhone   PhoneFormat
	MobilePhone PhoneFormat
}

type PhoneFormat struct {
	CountryCode string
	AreaCode    string
	Number      string
}

func Create(payload Client) (response string, err error) {

	payloadByte, payloadStringErr := json.Marshal(payload)

	if payloadStringErr != nil {
		return "", payloadStringErr
	}

	connection := utils.Connection{
		Url:     enums.BASE_URL + "/customers",
		Method:  "POST",
		Payload: string(payloadByte),
	}

	response, responseErr := utils.SendRequest(connection)

	if responseErr != nil {
		return "", responseErr
	}

	return response, nil
}
