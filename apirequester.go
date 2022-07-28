package oygo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// APIRequester abstraction of HTTP Client that will make API calls to OY backend.
// `body` is POST-requests' bodies if applicable.
// `result` pointer to value which response string will be unmarshalled to.
type APIRequester interface {
	Call(ctx context.Context, method string, url string, header *http.Header, body interface{}, result interface{}) *Error
}

// APIRequesterImplementation is the default implementation of APIRequester
type APIRequesterImplementation struct {
	HTTPClient *http.Client
}

// Call makes HTTP requests with JSON-format body.
// `body` is POST-requests' bodies if applicable.
// `result` pointer to value which response string will be unmarshalled to.
func (a *APIRequesterImplementation) Call(ctx context.Context, method string, url string, header *http.Header, body interface{}, result interface{}) *Error {
	reqBody := []byte("")
	var req *http.Request
	var err error

	isParamsNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())

	if !isParamsNil {
		reqBody, err = json.Marshal(body)
		log.Print(string(reqBody))
		if err != nil {
			return FromGoErr(err)
		}
	}

	req, err = newHTTPRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBody),
	)

	if err != nil {
		return FromGoErr(err)
	}

	if header != nil {
		req.Header = *header
	}
	req.Header.Set("Content-Type", "application/json")

	return a.doRequest(req, result)
}

func (a *APIRequesterImplementation) doRequest(req *http.Request, result interface{}) *Error {
	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return FromGoErr(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return FromGoErr(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return FromHTTPErr(resp.StatusCode, respBody)
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return FromGoErr(err)
	}

	return nil
}

func newHTTPRequestWithContext(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(
		ctx,
		method,
		url,
		body,
	)
}
