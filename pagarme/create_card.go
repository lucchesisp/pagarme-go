package pagarme

import (
	"context"
	"encoding/json"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/types"
)

func (i *Instance) CreateCard(ctx context.Context, clientId string, card *types.Card) (string, error) {
	payloadByte, err := json.Marshal(card)

	if err != nil {
		return "", err
	}

	connection := Connection{
		URL:       i.BaseURL + "/customers/" + clientId + "/cards",
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
