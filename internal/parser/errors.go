package parser

import "fmt"

// ErrInvalidInput indicates caller supplied invalid parameters.
type ErrInvalidInput string

func (e ErrInvalidInput) Error() string { return string(e) }

// ErrInvalidInputMsg builds an invalid-input error from a message.
func ErrInvalidInputMsg(msg string) error {
	return ErrInvalidInput(msg)
}

// InvalidInputf builds a formatted invalid-input error.
func InvalidInputf(format string, args ...any) error {
	return ErrInvalidInput(fmt.Sprintf(format, args...))
}
