package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnnaVyvert/safe-concept-server/cmd/server/routes"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/utils"
)

func init() {
	if err := utils.Load(); err != nil {
		log.Println(err)
	}
}

func serve() {
	port := utils.GetEnvDefault("PORT", "3000")
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("launching server on %v\n", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func main() {
	routes.DefineRoutes(http.DefaultServeMux)
	serve()
}
