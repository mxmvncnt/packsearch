package middleware

import "net/http"

func Combined(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return LoggingMiddleware(ErrorHandler(h))
}
