package middleware

import (
	"net/http"
	"time"

	"github.com/mxmvncnt/packsearch/server/utils/logger"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		crw := &customResponseWriter{ResponseWriter: w}

		logger.Infof(
			"[REQUEST START] Method: %s | Path: %s | IP: %s | User-Agent: %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			r.UserAgent(),
		)

		next.ServeHTTP(crw, r)

		duration := time.Since(start)
		logger.Infof(
			"[REQUEST END  ] Method: %s | Path: %s | Status: %d | Duration: %v | IP: %s | User-Agent: %s",
			r.Method,
			r.URL.Path,
			crw.statusCode,
			duration,
			r.RemoteAddr,
			r.UserAgent(),
		)
	}
}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
	errorSent  bool
}
