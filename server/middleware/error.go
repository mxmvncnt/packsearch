package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/mxmvncnt/packsearch/server/utils/apierror"
	"github.com/mxmvncnt/packsearch/server/utils/logger"
)

func ErrorHandler(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			switch err := err.(type) {
			case *apierror.ApiError:
				err.Send(w)
				logger.Errorf("[ErrorMiddleware] ApiError thrown: %v", *err)
			default:
				logger.Errorf("Unexpected exception caught: %s\n%v", formatRawValue(err), r)
				logger.Errorf("Stack Trace:\n%s\n", debug.Stack())
				apierror.NewApiError(
					http.StatusInternalServerError,
					"UNEXPECTED_ERROR",
					"Please wait and try again",
					"An unexpected internal error has occurred").Send(w)
			}
		}
	}
}

func formatRawValue(v interface{}) string {
	requestStr := fmt.Sprintf("%#v", v)
	requestStr = strings.ReplaceAll(requestStr, ", ", ",\n  ")
	requestStr = strings.ReplaceAll(requestStr, "{", "{\n  ")
	requestStr = strings.ReplaceAll(requestStr, "}", "\n}")
	requestStr = strings.ReplaceAll(requestStr, " map[", "\n  map[")
	requestStr = strings.ReplaceAll(requestStr, "] map[", "]\n  map[")
	requestStr = strings.ReplaceAll(requestStr, " 0x", "\n  0x")
	requestStr = strings.ReplaceAll(requestStr, " <nil>", "\n  <nil>")

	return requestStr
}
