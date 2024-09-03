package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/isaacmirandacampos/dreamkoffee/pkg/auth"
)

const userIDKey = contextKey("userID")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		auth, err := auth.JwtValidate(r.Context(), &cookie.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, auth.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(ctx context.Context) (int32, error) {
	const userIDKey = contextKey("userID")
	id, ok := ctx.Value(userIDKey).(int32)
	if !ok {
		return 0, errors.New("userID not found in context")
	}
	return id, nil
}
