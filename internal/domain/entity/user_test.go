package entity_test

import (
	"testing"

	"github.com/isaacmirandacampos/dreamkoffee/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	t.Parallel()
	t.Run("SetPassword", func(t *testing.T) {
		t.Run("Should_set_a_password", func(t *testing.T) {
			user := entity.User{}
			err := user.SetPassword("password")
			if err != nil {
				t.Errorf("SetPassword() error = %v, want nil", err)
			}
			if !user.Password.Valid {
				t.Error("SetPassword() user.Password.Valid = false, want true")
			}
		})

		t.Run("Should_return_an_error_when_the_password_is_short", func(t *testing.T) {
			user := entity.User{}
			err := user.SetPassword("pass")
			if err == nil {
				t.Error("SetPassword() error = nil, want an error")
			}
			assert.Equal(t, err.Error(), "password must have at least 8 characters")
		})

		t.Run("Should_return_an_error_when_the_password_is_empty", func(t *testing.T) {
			user := entity.User{}
			err := user.SetPassword("")
			if err == nil {
				t.Error("SetPassword() error = nil, want an error")
			}
			assert.Equal(t, err.Error(), "password cannot be empty")
		})
	})

	t.Run("PasswordIsValid", func(t *testing.T) {
		t.Run("Should_return_true_when_the_password_is_valid", func(t *testing.T) {
			user := entity.User{}
			user.SetPassword("password")
			valid, err := user.PasswordIsValid("password")
			if err != nil {
				t.Errorf("PasswordIsValid() error = %v, want nil", err)
			}
			if !valid {
				t.Error("PasswordIsValid() valid = false, want true")
			}
		})

		t.Run("Should_return_false_when_the_password_is_invalid", func(t *testing.T) {
			user := entity.User{}
			user.SetPassword("password")
			valid, err := user.PasswordIsValid("invalid_password")
			if err != nil {
				t.Errorf("PasswordIsValid() error = %v, want nil", err)
			}
			if valid {
				t.Error("PasswordIsValid() valid = true, want false")
			}
		})

		t.Run("Should_return_an_error_when_the_password_is_not_setted", func(t *testing.T) {
			user := entity.User{}
			valid, err := user.PasswordIsValid("password")
			if err == nil {
				t.Error("PasswordIsValid() error = nil, want an error")
			}
			assert.Equal(t, err.Error(), "password is not setted")
			if valid {
				t.Error("PasswordIsValid() valid = true, want false")
			}
		})
	})
}
