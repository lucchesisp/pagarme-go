package pagarme

import (
	"context"
	"errors"
	customError "github.com/lucchesisp/pagarme-go/errors"
	"github.com/lucchesisp/pagarme-go/types"
	"testing"
)

type CreatePlanSubscriptionRequestMock struct {
	handleServiceFn func(ctx context.Context, connection Connection) (response string, err error)
}

func (m CreatePlanSubscriptionRequestMock) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	return m.handleServiceFn(ctx, connection)
}

func TestCreatePlanSubscriptionWithoutAuthorization(t *testing.T) {
	handleServiceMock := CreatePlanSubscriptionRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", errors.New("authorization has been denied for this request")
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.CreatePlanSubscription(context.Background(), &types.PlanSubscription{})

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}

	expectedError := &customError.Error{
		ErrorCode:    500,
		ErrorMessage: "authorization has been denied for this request",
	}

	if responseErr.Error() != expectedError.Error() {
		t.Error("Expected authorization has been denied for this request, got ", responseErr.Error())
	}
}

func TestCreatePlanSubscriptionWithSuccess(t *testing.T) {
	handleServiceMock := DeleteCardRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "{\"status\":\"success\"}", nil
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.CreatePlanSubscription(context.Background(), &types.PlanSubscription{})

	if responseErr != nil {
		t.Error("Expected no errors, got ", responseErr)
	}

	if response == "" {
		t.Error("Expected response, got empty")
	}
}
