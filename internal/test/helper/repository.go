package helper

import (
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/mocks"
)

func RepositoryMock(repo *mocks.MockRepository) domain.Repository {
	return repo
}
