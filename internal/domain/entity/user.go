package entity

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/isaacmirandacampos/dreamkoffee/pkg/password_hashing"
)

type User struct {
	ID        int32          `db:"id" json:"id"`
	FullName  string         `db:"full_name" json:"full_name"`
	Email     string         `db:"email" json:"email"`
	Password  sql.NullString `db:"password" json:"-"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}

func NewUser(user *User) (User, error) {
	if user.Email == "" {
		return User{}, errors.New("email cannot be empty")
	}

	if !strings.Contains(user.Email, "@") {
		return User{}, errors.New("invalid email")
	}

	if user.FullName == "" {
		return User{}, errors.New("full name cannot be empty")
	}

	return User{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, nil
}

func (u *User) SetPassword(password string) error {
	err := u.acceptablePassword(password)
	if err != nil {
		return err
	}
	hash, err := password_hashing.Encrypt(password)
	if err != nil {
		return err
	}
	u.Password = sql.NullString{String: hash, Valid: true}
	return nil
}

func (u *User) PasswordIsValid(password string) (bool, error) {
	if !u.Password.Valid {
		return false, errors.New("password is not setted")
	}
	err := u.acceptablePassword(password)
	if err != nil {
		return false, err
	}
	equal, err := password_hashing.ComparePassword(password, u.Password.String)
	if err != nil {
		return false, err
	}
	return equal, nil
}

func (u *User) acceptablePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	if len(password) < 8 {
		return errors.New("password must have at least 8 characters")
	}

	return nil
}
