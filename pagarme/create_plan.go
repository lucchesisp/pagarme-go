package pagarme

import (
	"context"
	"encoding/json"
	"github.com/lucchesisp/pagarme-go/enums/method"
	"github.com/lucchesisp/pagarme-go/errors"
	"github.com/lucchesisp/pagarme-go/types"
)

func (i *Instance) CreatePlan(ctx context.Context, plan *types.Plan) (string, error) {
	payloadByte, _ := json.Marshal(plan)

	connection := Connection{
		URL:       i.BaseURL + "/plans",
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
