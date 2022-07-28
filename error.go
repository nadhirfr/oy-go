package oygo

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message,omitempty"`
	Url     string `json:"url,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// Error returns error message.
// This enables oygo.Error to comply with Go error interface
func (e *Error) Error() string {
	return e.Message
}

// GetStatus returns http status code
func (e *Error) GetStatus() int {
	return e.Status
}

// FromGoErr generates oygo.Error from generic go errors
func FromGoErr(err error) *Error {
	return &Error{
		Status:  http.StatusTeapot,
		Message: err.Error(),
		Url:     "",
	}
}

// FromHTTPErr generates oygo.Error from http errors with non 2xx status
func FromHTTPErr(status int, respBody []byte) *Error {
	var httpError *Error
	if err := json.Unmarshal(respBody, &httpError); err != nil {
		return FromGoErr(err)
	}
	httpError.Status = status

	return httpError
}
