package routes

import (
	"log"
	"net/http"

	"github.com/AnnaVyvert/safe-concept-server/cmd/server/handlers/file-exchanger/get_file"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/handlers/file-exchanger/put_file"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/handlers/root"
)

func Log(msg string, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(msg)
		handler(w, r)
	}
}

func DefineRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", Log("root handler", root.Handler))
	mux.HandleFunc("/get_file", Log("get_file handler", get_file.Handler))
	mux.HandleFunc("/put_file", Log("put_file handler", put_file.Handler))
}
