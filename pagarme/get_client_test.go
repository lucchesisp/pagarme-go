package pagarme

import (
	"context"
	"errors"
	customError "github.com/lucchesisp/pagarme-go/errors"
	"testing"
)

type GetClientRequestMock struct {
	handleServiceFn func(ctx context.Context, connection Connection) (response string, err error)
}

func (m GetClientRequestMock) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	return m.handleServiceFn(ctx, connection)
}

func TestGetClientWithoutAuthorization(t *testing.T) {
	handleServiceMock := GetClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", errors.New("authorization has been denied for this request")
	}

	HandleService = handleServiceMock

	clientID := "cus_eOP4preImI5V2G5K"
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetClient(context.Background(), clientID)

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

func TestGetClientWithSuccess(t *testing.T) {
	handleServiceMock := GetClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "{\"id\": \"cus_eOP4preImI5V2G5K\"}", nil
	}

	HandleService = handleServiceMock

	clientID := "cus_eOP4preImI5V2G5K"
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetClient(context.Background(), clientID)

	if responseErr != nil {
		t.Error("Expected nil, got ", responseErr)
	}

	if response != "{\"id\": \"cus_eOP4preImI5V2G5K\"}" {
		t.Error("Expected {\"id\": \"cus_eOP4preImI5V2G5K\"}, got ", response)
	}
}
