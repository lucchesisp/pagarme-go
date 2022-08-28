package pagarme

import (
	"context"
	"encoding/json"
	"errors"
	customError "github.com/lucchesisp/pagarme-go/errors"
	"github.com/lucchesisp/pagarme-go/types"
	"testing"
)

type EditClientRequestMock struct {
	handleServiceFn func(ctx context.Context, connection Connection) (response string, err error)
}

func (m EditClientRequestMock) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	return m.handleServiceFn(ctx, connection)
}

func TestEditClientWithoutAuthorization(t *testing.T) {
	handleServiceMock := EditClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", errors.New("authorization has been denied for this request")
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	clientID := "cus_eOP4preImI5V2G5K"

	response, responseErr := pagarme.EditClient(context.Background(), clientID, &types.Client{})

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}

	expectedError := customError.Error{
		ErrorCode:    500,
		ErrorMessage: "authorization has been denied for this request",
	}

	if responseErr.Error() != expectedError.Error() {
		t.Error("Expected errors, got ", responseErr)
	}
}

func TestEditClientWithoutClientID(t *testing.T) {
	handleServiceMock := EditClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", nil
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	clientID := ""

	response, responseErr := pagarme.EditClient(context.Background(), clientID, &types.Client{})

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}

	expectedError := customError.Error{
		ErrorCode:    400,
		ErrorMessage: customError.CustumerIDRequired,
	}

	if responseErr.Error() != expectedError.Error() {
		t.Error("Expected errors, got ", responseErr)
	}
}

func TestEditClientSuccess(t *testing.T) {
	handleServiceMock := EditClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "{\"id\": \"cus_eOP4preImI5V2G5K\"}", nil
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	clientID := "cus_eOP4preImI5V2G5K"

	response, responseErr := pagarme.EditClient(context.Background(), clientID, &types.Client{})

	if responseErr != nil {
		t.Error("Expected nil, got ", responseErr)
	}

	if response == "" {
		t.Error("Expected response, got empty")
	}

	var jsonMessage json.RawMessage
	jsonErr := json.Unmarshal([]byte(response), &jsonMessage)

	if jsonErr != nil {
		t.Error("Expected nil, got ", jsonErr)
	}
}
