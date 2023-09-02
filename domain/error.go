package domain

type InternalError struct{}

func NewInternalError() error {
	return InternalError{}
}

func (e InternalError) Error() string {
	return "Internal Error"
}

type ValidationError struct{}

func NewValidationError() error {
	return ValidationError{}
}

func (e ValidationError) Error() string {
	return "Validation Error"
}
