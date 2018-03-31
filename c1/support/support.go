package support

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var InternalServerErrorResponse = string(CreateResponseJSON(nil, []error{ErrInternalServer}))

func CaptureError(err error) {
	// Here is a good place to integrate error tracking libraries
}

func HandleJSONDecode(request interface{}, w http.ResponseWriter, r *http.Request) (success bool) {
	if r.ContentLength < 1 {
		respJ := CreateResponseJSON([]error{ErrBodyExpected}, nil)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respJ)
		return false
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(bytes.NewReader(bodyBytes))
	if decodeErr := decoder.Decode(request); decodeErr != nil {
		HandleJSONDecodeError(decodeErr, w)
		return false
	}

	return true
}

func HandleJSONDecodeError(err error, w http.ResponseWriter) {
	// Return 400 if it's an expected Decode error
	switch err.(type) {
	case *json.UnmarshalFieldError, *json.UnmarshalTypeError, *json.InvalidUnmarshalError:
		jsonResponse := CreateResponseJSON(nil, []error{err})

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
		CaptureError(err)

		return

	case *json.SyntaxError:
		jsonResponse := CreateResponseJSON(nil, []error{fmt.Errorf("Invalid JSON for request body")})

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
		CaptureError(err)

		return

	default:
		HandleInternalServerError(err, w)
	}
}

func HandleInternalServerError(err error, w http.ResponseWriter) {
	CaptureError(err)
	http.Error(w, InternalServerErrorResponse, http.StatusInternalServerError)
}

func HandleError(err error, w http.ResponseWriter) {
	if httpErr, ok := errors.Cause(err).(HTTPError); ok {
		respJ := CreateResponseJSON(nil, httpErr.HTTPResponse())

		w.WriteHeader(httpErr.HTTPCode())
		w.Write(respJ)
	} else {
		HandleInternalServerError(err, w)
	}
}
