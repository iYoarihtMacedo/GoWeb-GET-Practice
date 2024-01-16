package response

import "net/http"

func Text(w http.ResponseWriter, code int, body string) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(code)

	w.Write([]byte(body))

}
