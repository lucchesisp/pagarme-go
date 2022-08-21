package pagarme

import (
	"context"
	"fmt"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/errors"
)

// GetAllClients returns all clients with pagination
func (i *Instance) GetAllClients(ctx context.Context, page uint64, size uint64) (string, error) {

	if page == 0 || size == 0 {
		return "", &errors.Error{
			ErrorCode:    400,
			ErrorMessage: errors.PageAndSizeRequired,
		}
	}

	connection := Connection{
		URL:       i.BaseURL + fmt.Sprintf("/customers?page=%d&size=%d", page, size),
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
