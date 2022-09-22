package pagarme

import (
	"context"
	"errors"
	customError "github.com/lucchesisp/pagarme-go/errors"
	"github.com/lucchesisp/pagarme-go/types"
	"testing"
)

type CreateCardRequestMock struct {
	handleServiceFn func(ctx context.Context, connection Connection) (response string, err error)
}

func (m CreateCardRequestMock) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	return m.handleServiceFn(ctx, connection)
}

func TestCreateNewCardWithoutAuthorization(t *testing.T) {
	handleServiceMock := CreateCardRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", errors.New("authorization has been denied for this request")
	}

	HandleService = handleServiceMock

	clientID := "cus_eOP4preImI5V2G5K"
	card := &types.Card{}
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.CreateCard(context.Background(), clientID, card)

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}
}

func TestCreateNewCardWithoutClientID(t *testing.T) {
	handleServiceMock := CreateCardRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", nil
	}

	HandleService = handleServiceMock

	clientID := ""
	card := &types.Card{}
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.CreateCard(context.Background(), clientID, card)

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}

	expectedError := customError.Error{
		ErrorCode:    400,
		ErrorMessage: customError.CustomerIDRequired,
	}

	if responseErr.Error() != expectedError.Error() {
		t.Error("Expected errors, got ", responseErr)
	}
}

func TestCreateCardWithSuccess(t *testing.T) {
	handleServiceMock := CreateCardRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "{\"status\":\"success\"}", nil
	}

	HandleService = handleServiceMock

	clientID := "cus_eOP4preImI5V2G5K"
	card := &types.Card{
		CustomerID: clientID,
	}

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.CreateCard(context.Background(), clientID, card)

	if responseErr != nil {
		t.Error("Expected no errors, got ", responseErr)
	}

	if response == "" {
		t.Error("Expected response, got empty")
	}
}
