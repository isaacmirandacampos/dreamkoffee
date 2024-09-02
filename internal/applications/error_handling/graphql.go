package error_handling

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Graphql(ctx context.Context, status_code int, err string, message string, extra ...interface{}) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: message,
		Extensions: map[string]interface{}{
			"status_code": status_code,
			"error":       strings.ToUpper(err),
			"extra":       extra,
		},
	}
}
