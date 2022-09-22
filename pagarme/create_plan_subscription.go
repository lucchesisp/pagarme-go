package pagarme

import (
	"context"
	"encoding/json"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/errors"
	"github.com/lucchesisp/pagarme-go/types"
)

func (i *Instance) CreatePlanSubscription(ctx context.Context, plan *types.PlanSubscription) (string, error) {
	payloadByte, _ := json.Marshal(plan)

	connection := Connection{
		URL:       i.BaseURL + "/subscriptions",
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
