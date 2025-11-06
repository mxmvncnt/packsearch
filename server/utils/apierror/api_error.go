package apierror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	StatusCode    int    `json:"-"`
	Code          string `json:"code"`
	Message       string `json:"message"`
	Reason        string `json:"reason"`
	InternalError error  `json:"-"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf(
		"\n----> ApiError(\n\tCode: %s \n\tMessage: %s \n\tReason: %s\n\tInternal reason: %s\n      )",
		e.Code,
		e.Message,
		e.Reason,
		e.InternalError.Error(),
	)
}

func NewApiError(statusCode int, code string, message string, reason string) *ApiError {
	return &ApiError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Reason:     reason,
	}
}

func NewApiErrorWithError(statusCode int, code string, message string, reason string, err error) *ApiError {
	return &ApiError{
		StatusCode:    statusCode,
		Code:          code,
		Message:       message,
		Reason:        reason,
		InternalError: err,
	}
}

func SendApiError(w http.ResponseWriter, apiError *ApiError) {
	w.WriteHeader(apiError.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiError)
}

func (e *ApiError) Send(w http.ResponseWriter) {
	SendApiError(w, e)
}
