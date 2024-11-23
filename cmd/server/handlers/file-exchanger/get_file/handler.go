package get_file

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnnaVyvert/safe-concept-server/cmd/server/guards"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/vars"
	"github.com/AnnaVyvert/safe-concept-server/internal/datastore"
)

func isRequestHeadersValid(r *http.Request) bool {
	// db_header, ...
	return r.Method == "GET"
}

var _ datastore.Identifier = new(Id)

type Id string

// Identity implements datastore.Identifier.
func (i Id) Identity() string {
	return string(i)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if !isRequestHeadersValid(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientToken := r.Header.Get("token")

	if !guards.IsClientAuthorized(clientToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	storage := datastore.Crypted(datastore.NewFileStorage(vars.FsFolderPath))

	id := Id(getFileNameByToken(clientToken))
	file, err := storage.Load(id)
	if err != nil {
		// w.Write()
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file))
	if _, err := w.Write(file); err != nil {
		log.Println("can not write file:", err)
	}

	w.WriteHeader(http.StatusOK)
}
