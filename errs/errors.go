package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

// // Error implements error.
// func (*AppError) Error() string {
// 	panic("unimplemented")
// }

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}

}
func NewUnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}

}
