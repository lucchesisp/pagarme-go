package pagarme

import (
	"context"
	json2 "encoding/json"
	"errors"
	"github.com/lucchesisp/pagarme-go/types"
	"testing"
)

type SendRequestMock struct {
	handleServiceFn func(ctx context.Context, connection Connection) (response string, err error)
}

func (m SendRequestMock) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	return m.handleServiceFn(ctx, connection)
}

func TestCreateNewClientWithoutAuthorization(t *testing.T) {
	handleServiceMock := SendRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "", errors.New("authorization has been denied for this request")
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.CreateNewClient(context.Background(), &types.Client{})

	if responseErr == nil {
		t.Error("Expected error, got nil")
	}

	if response != "" {
		t.Error("Expected empty response, got ", response)
	}
}

func TestCreateNewClientSuccess(t *testing.T) {
	handleServiceMock := SendRequestMock{}
	handleServiceMock.handleServiceFn = func(ctx context.Context, connection Connection) (response string, err error) {
		return "{\"id\": \"cus_eOP4preImI5V2G5K\"}", nil
	}

	HandleService = handleServiceMock

	secretKey := "secretKey"
	pagarme, _ := Dial(secretKey)

	response, responseErr := pagarme.CreateNewClient(context.Background(), &types.Client{})

	if responseErr != nil {
		t.Error("Expected nil, got ", responseErr)
	}

	if response == "" {
		t.Error("Expected response, got empty")
	}

	var json json2.RawMessage
	jsonErr := json2.Unmarshal([]byte(response), &json)

	if jsonErr != nil {
		t.Error("Expected nil, got ", jsonErr)
	}
}
