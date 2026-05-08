package dto

import "go-backend/internal/common/pagination"

type ArticleCreateReq struct {
	Title    string  `json:"title" binding:"required"`
	Content  *string `json:"content" binding:"omitempty"`
	ImageUrl *string `json:"image_url" binding:"omitempty"`
}

type ArticleFindAllInput struct {
	pagination.Query
}
