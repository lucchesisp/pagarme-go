package pagarme

import (
	"context"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/errors"
)

// GetClient get a client entity by client id.
func (i *Instance) GetClient(ctx context.Context, id string) (string, error) {
	connection := Connection{
		URL:       i.BaseURL + "/customers/" + id,
		Method:    method.GET,
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
