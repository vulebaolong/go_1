package repository_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/ent/users"
	"go-backend/internal/dto"
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

// CreateUser implements [repository.UserRepository].
func (a *userRepository) CreateUser(ctx context.Context, body dto.AuthRegisterReq) (*ent.Users, error) {
	entCreate := a.entClient.Users.Create()
	entCreate = entCreate.SetEmail(body.Email)
	entCreate = entCreate.SetPassword(body.Password)
	entCreate = entCreate.SetFullName(body.FullName)

	return entCreate.Save(ctx)
}

// FindUserByEmail implements [repository.UserRepository].
func (a *userRepository) FindUserByEmail(ctx context.Context, email string) (*ent.Users, error) {
	entQuery := a.entClient.Users.Query()
	entQuery = entQuery.Where(users.EmailEQ(email))
	return entQuery.Only(ctx)
}

// FindUserById implements [repository.UserRepository].
func (a *userRepository) FindUserById(ctx context.Context, id int) (*ent.Users, error) {
	entQuery := a.entClient.Users.Query()
	entQuery = entQuery.Where(users.IDEQ(id))
	return entQuery.Only(ctx)
}
