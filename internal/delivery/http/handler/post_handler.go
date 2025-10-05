package handler

import (
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