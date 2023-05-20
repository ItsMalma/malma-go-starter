package exception

import (
	"fmt"
)

type ControllerError struct {
	Status  string
	Message string
	Code    int
}

func NewControllerError(status string, message string, code int) ControllerError {
	return ControllerError{
		Status:  status,
		Message: message,
		Code:    code,
	}
}

func (err ControllerError) Error() string {
	return fmt.Sprintf("%v %v: %v", err.Status, err.Code, err.Message)
}

func ErrParseRequest(reason string) ControllerError {
	return ControllerError{
		Status:  "UNPARSED_REQUEST",
		Message: fmt.Sprintf("Unable to parse request (%v)", reason),
		Code:    422,
	}
}

func ErrContentType(expected string, got string) ControllerError {
	return ControllerError{
		Status:  "UNSUPPORTED_CONTENT_TYPE",
		Message: fmt.Sprintf("Expected %v content type but got %v instead", expected, got),
		Code:    415,
	}
}
