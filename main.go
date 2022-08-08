package main

import (
	"fmt"
	"github.com/lucchesisp/pagarme-go/client"
	"time"
)

func main() {

	payload := client.Client{
		Name:           "Gabriel Lucchesi",
		Email:          "lucchesisp@gmail.com",
		Code:           "001",
		DocumentNumber: "05200645182",
		DocumentType:   "CPF",
		ClientType:     "individual",
		Gender:         "male",
		Address: client.Address{
			Country: "Brazil",
			State:   "São Paulo",
			City:    "São Paulo",
			ZipCode: "01308040",
		},
		Phones: []client.Phone{
			{
				MobilePhone: client.PhoneFormat{
					CountryCode: "55",
					AreaCode:    "11",
					Number:      "993417695",
				},
			},
		},
		Birthday: time.Now(),
		Metadata: "",
	}

	response, responseErr := client.Create(payload)

	if responseErr != nil {
		fmt.Errorf("Error: %s", responseErr)
	}

	fmt.Println(response)
}
