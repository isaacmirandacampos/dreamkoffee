package user

import (
	"context"

	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/error_handling"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/handlers/graph/model"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain/entity"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
)

func (u *userUseCase) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	userEntity, err := entity.NewUser(&entity.User{
		Email:    input.Email,
		FullName: input.FullName,
	})
	if err != nil {
		return nil, error_handling.Graphql(ctx, 400, "invalid_input", err.Error())
	}
	err = userEntity.SetPassword(input.Password)
	if err != nil {
		return nil, error_handling.Graphql(ctx, 400, "invalid_input", err.Error())
	}

	exists, err := u.repo.ExistsAnUserUsingTheSameEmail(ctx, userEntity.Email)
	if err != nil {
		return nil, error_handling.Graphql(ctx, 500, "internal_server_error", "Failed check if email already in use", err.Error())
	}

	if exists {
		return nil, error_handling.Graphql(ctx, 422, "email_in_use", "Email already in use")
	}

	user, err := u.repo.CreateUser(ctx, &persistence.CreateUserParams{
		FullName: userEntity.FullName,
		Email:    userEntity.Email,
		Password: userEntity.Password,
	})

	if err != nil {
		return nil, error_handling.Graphql(ctx, 401, "email_or_password_invalid", "Email or password invalid")
	}

	return &model.User{
		ID:        int(user.ID),
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
