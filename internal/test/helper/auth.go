package helper

import "context"

type contextKey string

func Auth(ctx context.Context, id int32) context.Context {
	const userIDKey = contextKey("userID")
	newCtx := context.WithValue(ctx, userIDKey, id)
	return newCtx
}
