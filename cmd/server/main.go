package main

import (
	"fmt"
	"log"
	"net/http"

	lib "github.com/AnnaVyvert/safe-concept-server/cmd/server/common_lib"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/routes"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/vars"
)

func init() {
	lib.PanicIfError(vars.Load())
}

func serve() {
	port := lib.GetEnv("PORT")
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("launching server on %v\n", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}


func main() {
	routes.DefineRoutes()
	serve()
}
