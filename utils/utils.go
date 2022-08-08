package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Connection struct {
	Url     string
	Payload string
	Method  string
}

func SendRequest(connection Connection) (response string, err error) {
	ioPayload := strings.NewReader(connection.Payload)
	req, reqErr := http.NewRequest(connection.Method, connection.Url, ioPayload)

	if reqErr != nil {
		return "", reqErr
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

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
