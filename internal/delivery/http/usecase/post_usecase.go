package usecase

import (
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/repository"
	"github.com/evandrarf/porto-ilits-backend/internal/domain"
	"github.com/evandrarf/porto-ilits-backend/internal/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	PostUsecase interface {
		Create(ctx *gin.Context, req domain.PostCreateRequest) (error)
		GetAll(ctx *gin.Context) ([]entity.Post, error)
	}

	postUsecase struct{
		db *gorm.DB
		postRepository repository.PostRepository
	}
)

func NewPostUsecase(db *gorm.DB, postRepository repository.PostRepository) PostUsecase {
	return &postUsecase{
		db: db,
		postRepository: postRepository,
	}
}

func (u *postUsecase) Create(ctx *gin.Context, req domain.PostCreateRequest) ( error) {
	tx := u.db.WithContext(ctx.Copy().Request.Context()).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return  tx.Error
	}

	data := entity.Post{
		Title: req.Title,
		Content: req.Content,
	}

	if err:= u.postRepository.Create(tx, &data); err != nil {
		tx.Rollback()
		return  err
	}

	if err := tx.Commit().Error; err != nil {
		return  err
	}

	return  nil
}

func (u *postUsecase) GetAll(ctx *gin.Context) ([]entity.Post, error) {
	var posts []entity.Post
	err := u.postRepository.All(u.db, &posts, nil, nil)
	if err != nil {
		return nil, err
	}

	return posts, nil
}