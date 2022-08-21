package pagarme

import (
	"context"
	"encoding/json"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/errors"
	"github.com/lucchesisp/pagarme-go/types"
)

// CreateCard create a new card.
func (i *Instance) CreateCard(ctx context.Context, clientID string, card *types.Card) (string, error) {

	if clientID == "" {
		return "", &errors.Error{
			ErrorCode:    400,
			ErrorMessage: errors.ClientIDRequired,
		}
	}

	payloadByte, _ := json.Marshal(card)

	connection := Connection{
		URL:       i.BaseURL + "/customers/" + clientID + "/cards",
		Method:    method.POST,
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
