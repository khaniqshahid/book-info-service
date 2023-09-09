package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// // Error implements error.
// func (*AppError) Error() string {
// 	panic("unimplemented")
// }

// Reciever Function to return only the error message and not the code by omitempty "Code" in json
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
func NewNotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}

}
func NewUnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}

}
