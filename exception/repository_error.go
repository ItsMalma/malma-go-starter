package exception

import "fmt"

type RepositoryErrorStatus string

const (
	RepositoryErrorStatusNotFound RepositoryErrorStatus = "NOT_FOUND"
)

type RepositoryError struct {
	Status  RepositoryErrorStatus
	Message string
}

func NewRepositoryError(status RepositoryErrorStatus, message string) RepositoryError {
	return RepositoryError{Status: status, Message: message}
}

func (err RepositoryError) Error() string {
	return fmt.Sprintf("%v: %v", err.Status, err.Message)
}
