package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, body any) {

	// check body
	if body == nil {
		w.WriteHeader(code)
		return
	}

	// marshal body
	bytes, err := json.Marshal(body)
	if err != nil {
		// default error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	w.Write(bytes)

}
