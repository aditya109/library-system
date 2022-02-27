package errors

import (
	"fmt"
)

// AppError is a custom error used within this application to raise custom exceptions
type AppError struct {
	Cause string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("cause: %s",
		e.Cause)
}
