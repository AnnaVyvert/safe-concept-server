package root

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "root")
}
