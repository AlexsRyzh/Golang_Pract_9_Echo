package error

type RequestError struct {
	StatusCode int
	Err        error
}
