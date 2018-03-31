package support

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
	HTTPCode() int
	HTTPResponse() []error
}

var ErrInternalServer = fmt.Errorf("Internal Server Error")
var ErrBodyExpected = fmt.Errorf("Expected Request Body")

type ErrDigestNotFound string

func (err ErrDigestNotFound) Error() string {
	return fmt.Sprintf("Digest %s not found", string(err))
}

func (err ErrDigestNotFound) HTTPCode() int {
	return http.StatusNotFound
}

func (err ErrDigestNotFound) HTTPResponse() []error {
	return []error{err}
}
