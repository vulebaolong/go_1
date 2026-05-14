package repository

import (
	"context"
)

type UserRepository interface {
	ExitsByEmail(ctx context.Context, email string) (bool, error)
}
