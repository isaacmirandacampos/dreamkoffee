package middleware

import (
	"context"
	"net/http"
)

const ResponseWriterKey = contextKey("httpResponseWriter")

func WithResponseWriter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ResponseWriterKey, w)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
