package repository_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/ent/users"
	"go-backend/internal/repository"
)

type userRepository struct {
	entClient *ent.Client
}

func NewUserRepository(entClient *ent.Client) repository.UserRepository {
	return &userRepository{
		entClient: entClient,
	}
}

// FindAll implements [repository.UserRepository].
func (a *userRepository) ExitsByEmail(ctx context.Context, email string) (bool, error) {
	entQuery := a.entClient.Users.Query()

	entQuery = entQuery.Where(users.EmailEQ(email))

	return entQuery.Exist(ctx)
}
