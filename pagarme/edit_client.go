package pagarme

import (
	"context"
	"encoding/json"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/errors"
	"github.com/lucchesisp/pagarme-go/types"
)

// EditClient edit a client entity.
func (i *Instance) EditClient(ctx context.Context, clientID string, client *types.Client) (string, error) {

	if clientID == "" {
		return "", &errors.Error{
			ErrorCode:    400,
			ErrorMessage: errors.ClientIDRequired,
		}
	}

	payloadByte, err := json.Marshal(client)

	if err != nil {
		return "", &errors.Error{
			ErrorCode:    400,
			ErrorMessage: errors.InvalidJSON,
		}
	}

	connection := Connection{
		URL:       i.BaseURL + "/customers/" + clientID,
		Method:    method.PUT,
		Payload:   string(payloadByte),
		SecretKey: i.SecretKey,
	}

	response, responseErr := HandleService.SendRequest(ctx, connection)

	if responseErr != nil {
		return "", &errors.Error{
			ErrorCode:    500,
			ErrorMessage: responseErr.Error(),
		}
	}

	return response, nil
}
