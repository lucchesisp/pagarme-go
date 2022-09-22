package pagarme

import (
	"context"
	"errors"
	customError "github.com/lucchesisp/pagarme-go/errors"
	"testing"
)

type GetAllClientsRequestMock struct {
	handleServiceFn func(ctx context.Context, connection Connection) (response string, err error)
}

func (m GetAllClientsRequestMock) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	return m.handleServiceFn(ctx, connection)
}

func TestGetAllClientsWithoutAuthorization(t *testing.T) {
	handleServiceMock := GetAllClientsRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", errors.New("authorization has been denied for this request")
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetAllClients(context.Background(), 1, 1)

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

func TestGetAllClientsWithZeroPage(t *testing.T) {
	handleServiceMock := GetAllClientsRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", nil
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetAllClients(context.Background(), 0, 0)

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}

	expectedError := customError.Error{
		ErrorCode:    400,
		ErrorMessage: customError.PageAndSizeRequired,
	}

	if responseErr.Error() != expectedError.Error() {
		t.Error("Expected errors, got ", responseErr)
	}
}

func TestGetAllClientsWithSuccess(t *testing.T) {
	handleServiceMock := GetAllClientsRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "{\"data\": [{\"id\": \"1\", \"name\": \"Lucca\"}]}", nil
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetAllClients(context.Background(), 1, 1)

	if responseErr != nil {
		t.Error("Expected no errors, got ", responseErr)
	}

	if response == "" {
		t.Error("Expected response, got empty")
	}
}
