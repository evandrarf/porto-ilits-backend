package handler

import (
	"fmt"
	"strconv"

	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/usecase"
	"github.com/evandrarf/porto-ilits-backend/internal/domain"
	"github.com/evandrarf/porto-ilits-backend/internal/pkg/response"
	"github.com/evandrarf/porto-ilits-backend/internal/pkg/validate"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	PostHandler interface {
		Create(ctx *gin.Context) 
		GetAll(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	postHandler struct{
		validator      *validate.Validator
		logger         *logrus.Logger
		postUsecase 		usecase.PostUsecase
	}
)

func NewPostHandler(validator *validate.Validator, logger *logrus.Logger, postUsecase usecase.PostUsecase) PostHandler {
	return &postHandler{
		validator:      validator,
		logger:         logger,
		postUsecase: postUsecase,
	}
}

func (h *postHandler) Create(ctx *gin.Context) {
	req := domain.PostCreateRequest{}

	if err := h.validator.ParseAndValidate(ctx, &req); err != nil {
		response.NewFailed(domain.CREATE_POST_FAILED, err, h.logger).Send(ctx)
		return
	}

	err := h.postUsecase.Create(ctx, req)
	if err != nil {
		response.NewFailed(domain.CREATE_POST_FAILED, err, h.logger).Send(ctx)
		return
	}

	response.NewSuccess(domain.CREATE_POST_SUCCESS, nil, nil).Send(ctx)
}

func (h *postHandler) GetAll(ctx *gin.Context) {
	posts, err := h.postUsecase.GetAll(ctx)
	if err != nil {
		response.NewFailed(domain.GET_POSTS_FAILED, err, h.logger).Send(ctx)
		return
	}

	response.NewSuccess(domain.GET_POSTS_SUCCESS, posts, nil).Send(ctx)
}

func (h *postHandler) Update(ctx *gin.Context) {
	req:= domain.PostUpdateRequest{}

	if err := h.validator.ParseAndValidate(ctx, &req); err != nil {
		response.NewFailed(domain.UPDATE_POST_FAILED, err, h.logger).Send(ctx)
		return
	}

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.NewFailed(domain.UPDATE_POST_FAILED, fmt.Errorf("invalid id parameter"), h.logger).Send(ctx)
		return
	}

	err = h.postUsecase.Update(ctx, uint(id), req)
	
	if err != nil {
		response.NewFailed(domain.UPDATE_POST_FAILED, err, h.logger).Send(ctx)
		return
	}
	
	response.NewSuccess(domain.UPDATE_POST_SUCCESS, nil, nil).Send(ctx)	
}

func (h *postHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.NewFailed(domain.DELETE_POST_FAILED, fmt.Errorf("invalid id parameter"), h.logger).Send(ctx)
		return
	}

	err = h.postUsecase.Delete(ctx, uint(id))
	if err != nil {
		response.NewFailed(domain.DELETE_POST_FAILED, err, h.logger).Send(ctx)
		return
	}

	response.NewSuccess(domain.DELETE_POST_SUCCESS, nil, nil).Send(ctx)
}