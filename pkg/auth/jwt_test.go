package auth_test

import (
	"context"
	"testing"

	"github.com/isaacmirandacampos/dreamkoffee/pkg/auth"
)

func TestJwt(t *testing.T) {
	ctx := context.Background()
	userID := int32(123)
	token, err := auth.JwtGenerate(ctx, &userID)
	if err != nil {
		t.Fatalf("error generating token: %v", err)
	}
	if token == "" {
		t.Fatalf("empty token")
	}
	t.Run("JwtValidate", func(t *testing.T) {
		t.Parallel()
		_, err := auth.JwtValidate(ctx, &token)
		if err != nil {
			t.Fatalf("error validating token: %v", err)
		}
	})

	t.Run("JwtValidateInvalidToken", func(t *testing.T) {
		t.Parallel()
		token := "invalidToken"
		_, err := auth.JwtValidate(ctx, &token)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
