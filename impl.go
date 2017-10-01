package httperr

func NewError(e error, statusCode int) Error {
	return ErrorImpl{
		error:      e,
		statusCode: statusCode,
	}
}

type ErrorImpl struct {
	error
	statusCode int
}

func (e ErrorImpl) HTTPStatusCode() int {
	return e.statusCode
}
