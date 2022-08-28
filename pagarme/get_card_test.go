package pagarme

import (
	"context"
	"errors"
	customError "github.com/lucchesisp/pagarme-go/errors"
	"testing"
)

type GetCardRequestMock struct {
	handleServiceFn func(ctx context.Context, connection Connection) (response string, err error)
}

func (m GetCardRequestMock) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	return m.handleServiceFn(ctx, connection)
}

func TestGetCardWithoutAuthorization(t *testing.T) {
	handleServiceMock := GetClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", errors.New("authorization has been denied for this request")
	}

	HandleService = handleServiceMock

	cardID := "123"
	customerID := "123"
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetCard(context.Background(), cardID, customerID)

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

func TestGetCardWithoutCardID(t *testing.T) {
	handleServiceMock := GetClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", nil
	}

	HandleService = handleServiceMock

	cardID := ""
	customerID := "123"
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetCard(context.Background(), cardID, customerID)

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}

	expectedError := &customError.Error{
		ErrorCode:    400,
		ErrorMessage: customError.CardIDRequired,
	}

	if responseErr.Error() != expectedError.Error() {
		t.Error("Expected CardIDRequired, got ", responseErr.Error())
	}
}

func TestGetCardWithoutCustomerID(t *testing.T) {
	handleServiceMock := GetClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", nil
	}

	HandleService = handleServiceMock

	cardID := "123"
	customerID := ""
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetCard(context.Background(), cardID, customerID)

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}

	expectedError := &customError.Error{
		ErrorCode:    400,
		ErrorMessage: customError.CustumerIDRequired,
	}

	if responseErr.Error() != expectedError.Error() {
		t.Error("Expected CustomerIDRequired, got ", responseErr.Error())
	}
}

func TestGetCardWithSuccess(t *testing.T) {
	handleServiceMock := GetClientRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "{\"status\":\"success\"}", nil
	}

	HandleService = handleServiceMock

	cardID := "123"
	customerID := "123"
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.GetCard(context.Background(), cardID, customerID)

	if responseErr != nil {
		t.Error("Expected no errors, got ", responseErr)
	}

	if response == "" {
		t.Error("Expected response, got empty")
	}
}
