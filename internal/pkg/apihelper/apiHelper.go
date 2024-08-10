package apihelper

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func Validate(s interface{}) error {
	return validate.Struct(s)
}

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return c.Message
}

func NotFound(message string) error {
	return &CustomError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func InternalError(message string) error {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func TooManyRequest(message string) error {
	return &CustomError{
		Code:    http.StatusTooManyRequests,
		Message: message,
	}
}

func BadRequest(message string) error {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func GromError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NotFound(err.Error())
	}

	return InternalError(err.Error())
}

func S3Error(err error) error {
	s3Err, ok := err.(awserr.Error)
	if !ok {
		return err
	}

	switch s3Err.Code() {
	case "NotFound":
		return NotFound(s3Err.Message())
	default:
		return InternalError(s3Err.Message())
	}
}

func ValidationError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(&map[string]any{
		"error": err.Error(),
	})
}

func Error(w http.ResponseWriter, err error) {
	customError := err.(*CustomError)
	w.WriteHeader(customError.Code)

	json.NewEncoder(w).Encode(&map[string]any{
		"error": customError,
	})
}

func Response(w http.ResponseWriter, data any) {
	json.NewEncoder(w).Encode(&map[string]any{
		"data": data,
	})
}
