package repository

import (
	"github.com/evandrarf/porto-ilits-backend/internal/entity"
	"gorm.io/gorm"
)

type (
	PostRepository interface {
		Repository[entity.Post]
	}

	postRepository struct{
		repository[entity.Post]
	}
)

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		repository: repository[entity.Post]{
			DB: db,
		},
	}
}