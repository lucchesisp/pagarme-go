package pagarme

import (
	"context"
	"encoding/json"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/types"
)

// CreateNewClient create a new client entity.
func (i *Instance) CreateNewClient(ctx context.Context, client *types.Client) (string, error) {
	payloadByte, _ := json.Marshal(client)

	connection := Connection{
		URL:       i.BaseURL + "/customers",
		Method:    method.POST,
		Payload:   string(payloadByte),
		SecretKey: i.SecretKey,
	}

	response, responseErr := HandleService.SendRequest(ctx, connection)

	if responseErr != nil {
		return "", responseErr
	}

	return response, nil
}
