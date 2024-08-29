package utils

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ErrorHandling(ctx context.Context, status_code int, err, message string) {
	graphql.AddError(ctx, &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: message,
		Extensions: map[string]interface{}{
			"status_code": status_code,
			"error":       err,
		},
	})
}
