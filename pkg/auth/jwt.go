package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaim struct {
	ID int32 `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "aSecret"
	}
	return secret
}

func JwtGenerate(ctx context.Context, userID *int32) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID: *userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtValidate(ctx context.Context, token *string) (*JwtCustomClaim, error) {
	t, err := jwt.ParseWithClaims(*token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	claims := &JwtCustomClaim{
		ID: t.Claims.(*JwtCustomClaim).ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t.Claims.(*JwtCustomClaim).ExpiresAt,
			IssuedAt:  t.Claims.(*JwtCustomClaim).IssuedAt,
		},
	}
	return claims, nil
}
