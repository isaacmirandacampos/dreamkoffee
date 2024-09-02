package auth

import (
	"net/http"
	"os"
	"time"
)

func Cookie(value string) *http.Cookie {
	return &http.Cookie{
		Name:     nameResolve(),
		Value:    value,
		HttpOnly: true,
		Domain:   domainResolve(),
		MaxAge:   time.Now().Add(168 * time.Hour).Second(),
	}
}

func domainResolve() string {
	environment := os.Getenv("ENV")
	if environment == "production" {
		return "https://dreamkoffee.com"
	}
	if environment == "staging" {
		return "https://staging.dreamkoffee.com"
	}
	return ""
}

func nameResolve() string {
	environment := os.Getenv("ENV")
	if environment == "staging" {
		return "auth_token_staging"
	}
	return "auth_token"
}
