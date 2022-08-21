package pagarme

import (
	"context"
	"github.com/lucchesisp/pagarme-go/enums/method"
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
		return "", responseErr
	}

	return response, nil
}
