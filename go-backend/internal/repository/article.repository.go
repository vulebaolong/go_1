package repository

import (
	"context"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
)

type ArticleRepository interface {
	Create(ctx context.Context, body dto.ArticleCreateReq) (any, error)
	CreateGorm(ctx context.Context, body dto.ArticleCreateReq) (any, error)
	GetAll(ctx context.Context, query pagination.Query) (any, error)
}
