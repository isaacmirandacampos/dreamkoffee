package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	userID := middleware.GetUserID(ctx)
	if userID == 0 {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
