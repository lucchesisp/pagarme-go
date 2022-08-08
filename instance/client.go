package instance

import (
	"context"
	"encoding/json"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/types"
)

// CreateNewClient create a new client entity.
func (i *Instance) CreateNewClient(ctx context.Context, client types.Client) (string, error) {
	payloadByte, payloadStringErr := json.Marshal(client)

	if payloadStringErr != nil {
		return "", payloadStringErr
	}

	connection := Connection{
		Url:       i.BaseUrl + "/customers",
		Method:    method.POST,
		Payload:   string(payloadByte),
		SecretKey: i.SecretKey,
	}

	response, responseErr := SendRequest(ctx, connection)

	if responseErr != nil {
		return "", responseErr
	}

	return response, nil
}

// EditClient edit a client entity.
func (i *Instance) EditClient(ctx context.Context, clientId string, client types.Client) (string, error) {
	payloadByte, payloadStringErr := json.Marshal(client)

	if payloadStringErr != nil {
		return "", payloadStringErr
	}

	connection := Connection{
		Url:       i.BaseUrl + "/customers/" + clientId,
		Method:    method.PUT,
		Payload:   string(payloadByte),
		SecretKey: i.SecretKey,
	}

	response, responseErr := SendRequest(ctx, connection)

	if responseErr != nil {
		return "", responseErr
	}

	return response, nil
}
