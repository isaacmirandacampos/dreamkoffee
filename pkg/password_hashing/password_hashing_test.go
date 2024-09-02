package password_hashing_test

import (
	"strings"
	"testing"

	"github.com/isaacmirandacampos/dreamkoffee/pkg/password_hashing"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()
	t.Run("should return a string with the encoded salt and hash", func(t *testing.T) {
		hash, err := password_hashing.Encrypt("password")
		if err != nil {
			t.Errorf("expected nil; got %v", err)
		}
		if hash == "" {
			t.Errorf("expected not empty; got empty")
		}
		assert.Equal(t, len(strings.Split(hash, ".")), 2)
	})
}
