package user

import (
	"context"
	"net/http"

	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/error_handling"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/handlers/graph/model"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/middleware"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain/entity"
	"github.com/isaacmirandacampos/dreamkoffee/pkg/auth"
)

func (u *userUseCase) Authenticate(ctx context.Context, email string, password string) (*model.AuthenticationResponse, error) {
	userExist, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return &model.AuthenticationResponse{
			Success: false,
		}, error_handling.Graphql(ctx, 401, "email_or_password_invalid", "Email or password invalid")
	}
	user, _ := entity.NewUser(&entity.User{
		ID:       userExist.ID,
		Email:    userExist.Email,
		Password: userExist.Password,
		FullName: userExist.FullName,
	})

	valid, err := user.PasswordIsValid(password)
	if err != nil || !valid {
		return &model.AuthenticationResponse{
			Success: false,
		}, error_handling.Graphql(ctx, 401, "email_or_password_invalid", "Email or password invalid")
	}

	w, ok := ctx.Value(middleware.ResponseWriterKey).(http.ResponseWriter)
	if !ok {
		return &model.AuthenticationResponse{
			Success: false,
		}, error_handling.Graphql(ctx, 500, "failed_set_cookie", "failed to retrieve response writer from context")
	}

	token, err := auth.JwtGenerate(ctx, &user.ID)
	if err != nil {
		return &model.AuthenticationResponse{
			Success: false,
		}, error_handling.Graphql(ctx, 500, "failed_generate_token", "failed to generate token")
	}

	http.SetCookie(w, auth.Cookie(token))

	return &model.AuthenticationResponse{
		Success: true,
	}, nil
}
