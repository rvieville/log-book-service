package apihelper

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

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
