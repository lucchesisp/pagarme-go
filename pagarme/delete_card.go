package pagarme

import (
	"context"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/errors"
)

// DeleteCard deletes the card with the given cardID and customerID.
func (i *Instance) DeleteCard(ctx context.Context, cardID string, customerID string) (string, error) {
	if cardID == "" {
		return "", &errors.Error{
			ErrorCode:    400,
			ErrorMessage: errors.CardIDRequired,
		}
	}

	if customerID == "" {
		return "", &errors.Error{
			ErrorCode:    400,
			ErrorMessage: errors.CustomerIDRequired,
		}
	}

	connection := Connection{
		URL:       i.BaseURL + "/customers/" + customerID + "/cards/" + cardID,
		Method:    method.DELETE,
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
