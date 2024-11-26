package put_file

import (
	"net/http"

	"github.com/AnnaVyvert/safe-concept-server/cmd/server/guards"
)

func isRequestHeadersValid(r *http.Request) bool {
	// db_header, ...
	return r.Method == "PUT"
}

func isFileValid() bool {
	// size, format, content
	return true
}

func getFileNameFromRequest(r *http.Request) string {
	return "mock_file_name"
}

func getFileContentFromRequest(r *http.Request) string {
	return "mock_file_content"
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if !isRequestHeadersValid(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := r.Header.Get("token")

	if !guards.IsClientAuthorized(token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !isFileValid() {
		// w.WriteHeader(http.StatusRequestEntityTooLarge)
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	saveFile(getFileNameFromRequest(r), getFileContentFromRequest(r))

	w.WriteHeader(http.StatusOK)
}
