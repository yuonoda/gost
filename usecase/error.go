package usecase

type InternalError struct{}

func NewInternalError() error {
	return InternalError{}
}

func (e InternalError) Error() string {
	return "Internal Error"
}

type ValidationError struct{}

func (e ValidationError) Error() string {
	return "Validation Error"
}
