package errors

import "fmt"

// Error This struct is used to represent an internal error.
type Error struct {
	ErrorCode    int
	ErrorMessage string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d]: %s", e.ErrorCode, e.ErrorMessage)
}
