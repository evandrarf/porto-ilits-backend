package validate

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidator() *Validator {
	validate := validator.New(validator.WithRequiredStructEnabled())

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{
		validate: validate,
		trans:    trans,
	}
}

func (v *Validator) ParseAndValidate(ctx *gin.Context, req interface{}) error {
	// Parse JSON body
	if err := ctx.ShouldBindJSON(req); err != nil {
		return errors.New("invalid request body")
	}

	err := v.validate.Struct(req)
	if err == nil {
		return nil
	}

	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors.New("request validation failed")
	}

	fields := v.translateError(validationErrs)
	return NewFieldsError(fields)
}

// translateError mengubah pesan error ke bahasa yang sudah diset
func (v *Validator) translateError(errs validator.ValidationErrors) map[string]string {
	fields := make(map[string]string)
	for _, e := range errs {
		fields[e.Field()] = e.Translate(v.trans)
	}
	return fields
}

type FieldsError struct {
	Fields map[string]string `json:"fields"`
}

func (e *FieldsError) Error() string {
	return "validation error"
}

func NewFieldsError(fields map[string]string) error {
	return &FieldsError{Fields: fields}
}
