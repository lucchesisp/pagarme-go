package pagarme

import (
	"context"
	"errors"
	customError "github.com/lucchesisp/pagarme-go/errors"
	"testing"
)

type DeleteCardRequestMock struct {
	handleServiceFn func(ctx context.Context, connection Connection) (response string, err error)
}

func (m DeleteCardRequestMock) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	return m.handleServiceFn(ctx, connection)
}

func TestDeleteCardWithoutAuthorization(t *testing.T) {
	handleServiceMock := DeleteCardRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", errors.New("authorization has been denied for this request")
	}

	HandleService = handleServiceMock

	cardID := "123"
	customerID := "123"
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.DeleteCard(context.Background(), cardID, customerID)

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

func TestDeleteCardWithoutCardID(t *testing.T) {
	handleServiceMock := DeleteCardRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", nil
	}

	HandleService = handleServiceMock

	cardID := ""
	customerID := "123"
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.DeleteCard(context.Background(), cardID, customerID)

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

func TestDeleteCardWithoutCustomerID(t *testing.T) {
	handleServiceMock := DeleteCardRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", nil
	}

	HandleService = handleServiceMock

	cardID := "123"
	customerID := ""
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.DeleteCard(context.Background(), cardID, customerID)

	if responseErr == nil {
		t.Error("Expected errors, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}

	expectedError := &customError.Error{
		ErrorCode:    400,
		ErrorMessage: customError.CustomerIDRequired,
	}

	if responseErr.Error() != expectedError.Error() {
		t.Error("Expected CustomerIDRequired, got ", responseErr.Error())
	}
}

func TestDeleteCardWithSuccess(t *testing.T) {
	handleServiceMock := DeleteCardRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "{\"status\":\"success\"}", nil
	}

	HandleService = handleServiceMock

	cardID := "123"
	customerID := "123"
	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.DeleteCard(context.Background(), cardID, customerID)

	if responseErr != nil {
		t.Error("Expected no errors, got ", responseErr)
	}

	if response == "" {
		t.Error("Expected response, got empty")
	}
}
