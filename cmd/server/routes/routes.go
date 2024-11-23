package routes

import (
	"net/http"

	"github.com/AnnaVyvert/safe-concept-server/cmd/server/handlers/file-exchanger/get_file"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/handlers/file-exchanger/put_file"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/handlers/root"
)

func DefineRoutes() {
	http.HandleFunc("/", root.Handler)
	http.HandleFunc("/get_file", get_file.Handler)
	http.HandleFunc("/put_file", put_file.Handler)
}
