<img src="https://avatars1.githubusercontent.com/u/3846050?v=4&s=200" width="127px" height="127px" align="left" margin="0 20px !important" />

# Pagar.me Golang
[![Github All Releases](https://img.shields.io/github/downloads/lucchesisp/pagarme-go/total.svg)]()
[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/lucchesisp/pagarme-go?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucchesisp/pagarme-go)](https://goreportcard.com/report/github.com/lucchesisp/pagarme-go)
[![Coverage Status](https://coveralls.io/repos/github/lucchesisp/pagarme-go/badge.svg?branch=master)](https://coveralls.io/github/lucchesisp/pagarme-go?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)


A Golang library to interact with Pagar.me API.

The documentation can be found in our [Go Reference](https://pkg.go.dev/github.com/lucchesisp/pagarme-go?tab=doc)

## Description

This library covers the following features:

[Clients]

- Create new client
- Edit client
- Get client
- Get all clients

[Cards]

- Create card
- Read card
- Delete card [Pending]

[Plans]

- Create plan [Pending]
- Edit plan [Pending]
- Delete plan [Pending]
- Get plan [Pending]
- Get all plans [Pending]

[Subscribe]

- Create subscribe [Pending]
- Create subscribe with plan [Pending]
- Edit subscribe [Pending]
- Cancel subscribe [Pending]


## How to use

```bash
go get -u github.com/lucchesisp/pagarme-go
```

## Example: Create new client

```go
pagarme, err := pagarme.Dial("YOUR_API_SECRET")
	
if err != nil {
  fmt.Println(err)
  return
}
	
client := types.Client{
  Code:           "#00000001",
  Name:           "Lucas Santos",
  Email:          "lucas.santos@gmail.com",
  DocumentNumber: "12345678909",
  DocumentType:   "CPF",
  Phones: types.Phone{
    MobilePhone: types.PhoneFormat{
      CountryCode: "55",
      AreaCode:    "11",
      Number:      "988888888",
    },
    HomePhone: types.PhoneFormat{
      CountryCode: "55",
      AreaCode:    "11",
      Number:      "988888888",
    },
  },
  Birthday: time.Now(),
  Address: types.Address{
    Country: "BR",
    State:   "SP",
    City:    "São Paulo",
    Line1:   "Rua Teste",
    ZipCode: "01311000",
  },
  ClientType: "individual",
  Gender:     "male",
  Metadata:   "",
}
	
response, err := pagarme.CreateNewClient(context.Background(), &client)

if err != nil {
  fmt.Println(err)
  return
}
	
fmt.Println(response)
```

If you support this project, consider buying me a coffee ;)

<a href="https://www.buymeacoffee.com/lucchesisp" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 50px !important;width: 150px !important;" ></a>
