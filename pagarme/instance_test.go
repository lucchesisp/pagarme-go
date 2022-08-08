package pagarme

import (
	"github.com/lucchesisp/pagarme-go/enums/config"
	"testing"
)

func TestDialWithoutSecretKey(t *testing.T) {
	secretKey := ""
	pagarme, pagarmeErr := Dial(secretKey)

	if pagarmeErr == nil {
		t.Error("Expected error, got nil")
	}

	if pagarme != nil {
		t.Error("Pagarme is not nil")
	}
}

func TestDialWithSuccess(t *testing.T) {
	secretKey := "secretKey"

	pagarme, pagarmeErr := Dial(secretKey)

	if pagarmeErr != nil {
		t.Error("Expected nil, got ", pagarmeErr)
	}

	if pagarme == nil {
		t.Error("Pagarme is nil")
	}

	if pagarme.SecretKey != secretKey {
		t.Error("SecretKey is not equal")
	}

	if pagarme.BaseUrl != config.BASE_URL {
		t.Error("BaseUrl is not equal")
	}
}