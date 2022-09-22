package pagarme

import (
	"context"
	"github.com/lucchesisp/pagarme-go/enums/config"
	"github.com/lucchesisp/pagarme-go/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Interface for the pagarme client
type Interface interface {
	SendRequest(ctx context.Context, connection Connection) (response string, err error)
}

// Impl is the implementation of the pagarme client
type Impl struct{}

// HandleService is the service that will handle the request
var HandleService Interface = &Impl{}

// Connection is the connection that will be used to send the request
type Connection struct {
	URL       string
	Payload   string
	Method    string
	SecretKey string
}

// Instance is the pagarme client
type Instance struct {
	Context   context.Context
	BaseURL   string
	SecretKey string
}

// Dial creates a new pagarme client
func Dial(secretKey string) (*Instance, error) {
	if len(secretKey) == 0 {
		return nil, &errors.Error{
			ErrorCode:    400,
			ErrorMessage: errors.SecretKeyRequired,
		}
	}

	return DialContext(context.Background(), secretKey), nil
}

// DialContext creates a new pagarme client with a context
func DialContext(ctx context.Context, secretKey string) *Instance {
	return &Instance{
		Context:   ctx,
		BaseURL:   config.BaseURL,
		SecretKey: secretKey,
	}
}

func DialWithoutContext(secretKey string) *Instance {
	return &Instance{
		BaseURL:   config.BaseURL,
		SecretKey: secretKey,
	}
}

// SendRequest sends the request to the pagarme server
func (i Impl) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	ioPayload := strings.NewReader(connection.Payload)
	req, reqErr := http.NewRequestWithContext(ctx, connection.Method, connection.URL, ioPayload)

	if reqErr != nil {
		return "", reqErr
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	req.SetBasicAuth(connection.SecretKey, "")

	res, resErr := http.DefaultClient.Do(req)

	if resErr != nil {
		return "", resErr
	}

	defer res.Body.Close()

	body, bodyErr := ioutil.ReadAll(res.Body)

	if bodyErr != nil {
		return "", bodyErr
	}

	return string(body), nil
}
