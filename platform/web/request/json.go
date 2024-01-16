package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrRequestContentTypeNotJSON = errors.New("request content type is not application/json")
	ErrRequestJSONInvalid        = errors.New("request json invalid")
)

// JSON decodes json from request body to pth
func JSON(r *http.Request, ptr any) (err error) {

	// check content type

	if r.Header.Get("Content-Type") != "application/json" {
		return ErrRequestContentTypeNotJSON

	}

	// get body
	err = json.NewDecoder(r.Body).Decode(ptr)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrRequestJSONInvalid, err)
		return

	}

	return
}
