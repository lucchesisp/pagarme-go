package pagarme

import (
	"context"
	"fmt"
	"github.com/lucchesisp/pagarme-go/enums/method"
)

// GetAllClients returns all clients with pagination
func (i *Instance) GetAllClients(ctx context.Context, page uint64, size uint64) (string, error) {
	connection := Connection{
		URL:       i.BaseURL + fmt.Sprintf("/customers?page=%d&size=%d", page, size),
		Method:    method.GET,
		SecretKey: i.SecretKey,
	}

	response, responseErr := HandleService.SendRequest(ctx, connection)

	if responseErr != nil {
		return "", responseErr
	}

	return response, nil
}
