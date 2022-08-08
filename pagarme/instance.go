package pagarme

import (
	"context"
	"fmt"
	"github.com/lucchesisp/pagarme-go/enums/config"
	"io/ioutil"
	"net/http"
	"strings"
)

type Interface interface {
	SendRequest(ctx context.Context, connection Connection) (response string, err error)
}

type Impl struct{}

var HandleService Interface = &Impl{}

type Connection struct {
	Url       string
	Payload   string
	Method    string
	SecretKey string
}

type Instance struct {
	Context   context.Context
	BaseUrl   string
	SecretKey string
}

func Dial(secretKey string) (*Instance, error) {
	if len(secretKey) == 0 {
		return nil, fmt.Errorf("secretKey is empty")
	}

	return DialContext(context.Background(), secretKey), nil
}

func DialContext(ctx context.Context, secretKey string) *Instance {
	return &Instance{
		Context:   ctx,
		BaseUrl:   config.BASE_URL,
		SecretKey: secretKey,
	}
}

func (i Impl) SendRequest(ctx context.Context, connection Connection) (response string, err error) {
	ioPayload := strings.NewReader(connection.Payload)
	req, reqErr := http.NewRequestWithContext(ctx, connection.Method, connection.Url, ioPayload)

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
