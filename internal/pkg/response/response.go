package response

import (
	"net/http"

	"github.com/evandrarf/porto-ilits-backend/internal/pkg/validate"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	StatusCode int    `json:"-"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
	Error      any    `json:"error,omitempty"`
	Data       any    `json:"data,omitempty"`
	Meta       any    `json:"meta,omitempty"`
}

func NewInternalServerError() *Response {
	return &Response{
		Success:    false,
		Message:    "Internal Server Error",
		StatusCode: http.StatusInternalServerError,
	}
}

func NewFailed(msg string, err error, logger *logrus.Logger) *Response {
	res := &Response{
		Success:    false,
		Message:    msg,
		StatusCode: http.StatusInternalServerError,
	}

	switch e := err.(type) {
	case *validate.FieldsError:
		res.StatusCode = http.StatusBadRequest
		res.Error = e.Fields
	default:
		if err != nil {
			res.Error = err.Error()
		}
	}

	if logger != nil && res.StatusCode >= http.StatusInternalServerError {
		logger.Error(err)
	}

	return res
}

func NewSuccess(msg string, data any, meta any) *Response {
	return &Response{
		Success: true,
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r)
}
