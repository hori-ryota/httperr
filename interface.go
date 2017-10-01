// httperr is interface package for http error
package httperr

//go:generate mockgen -source interface.go -package httperr -destination mock.go
type Error interface {
	Error() string
	HTTPStatusCode() int
}
