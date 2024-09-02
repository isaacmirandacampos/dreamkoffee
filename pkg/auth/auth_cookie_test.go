package auth_test

import (
	"testing"
	"time"

	"github.com/isaacmirandacampos/dreamkoffee/pkg/auth"
	"gotest.tools/assert"
)

func TestAuthCookie(t *testing.T) {
	t.Run("Cookie", func(t *testing.T) {
		t.Parallel()
		cookie := auth.Cookie("value")
		if cookie.Name != "auth_token" {
			t.Errorf("expected name: auth_token, got: %s", cookie.Name)
		}
		if cookie.Value != "value" {
			t.Errorf("expected value: value, got: %s", cookie.Value)
		}
		if !cookie.HttpOnly {
			t.Errorf("expected HttpOnly: true, got: %v", cookie.HttpOnly)
		}
		if cookie.Domain != "" {
			t.Errorf("expected Domain: \"\", got: %s", cookie.Domain)
		}
		assert.Equal(t, cookie.MaxAge, time.Now().Add(168*time.Hour).Second())
	})
}
