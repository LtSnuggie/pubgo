package pubgo

type InvalidKeyError struct {
	message string
}

func NewInvalidKeyError(url string) *InvalidKeyError {
	return &InvalidKeyError{
		message: "API key invalid or missing: " + url,
	}
}
func (e *InvalidKeyError) Error() string {
	return e.message
}

type NotFoundError struct {
	message string
}

func NewNotFoundError(url string) *NotFoundError {
	return &NotFoundError{
		message: "The specified resource was not found: " + url,
	}
}
func (e *NotFoundError) Error() string {
	return e.message
}

type IncorrectContentTypeError struct {
	message string
}

func NewIncorrectContentTypeError(url string) *IncorrectContentTypeError {
	return &IncorrectContentTypeError{
		message: "Content type incorrect or not specified: " + url,
	}
}
func (e *IncorrectContentTypeError) Error() string {
	return e.message
}

type TooManyRequestsError struct {
	message string
}

func NewTooManyRequestsError(url string) *TooManyRequestsError {
	return &TooManyRequestsError{
		message: "Too many requests: " + url,
	}
}
func (e *TooManyRequestsError) Error() string {
	return e.message
}

type UnhandledStatusCodeError struct {
	message string
}

func NewUnhandledStatusCodeError(url, status string) *UnhandledStatusCodeError {
	return &UnhandledStatusCodeError{
		message: "Unable to handle status code: " + status + " - " + url,
	}
}
func (e *UnhandledStatusCodeError) Error() string {
	return e.message
}
