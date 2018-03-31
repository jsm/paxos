package support

import "encoding/json"

type errorResponse struct {
	Errors []string `json:"errors"`
}

func CreateResponseJSON(data interface{}, errors []error) []byte {
	respJ := data

	if len(errors) > 0 {
		errStrings := []string{}

		for _, err := range errors {
			errStrings = append(errStrings, err.Error())
		}

		respJ = errorResponse{
			Errors: errStrings,
		}
	}

	respBytes, err := json.Marshal(respJ)
	if err != nil {
		panic(err)
	}
	return respBytes
}
