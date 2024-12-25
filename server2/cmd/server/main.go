package main

import (
	"log"
	"net/http"
	"os"

	api "server2/internal"
)

func getPort() string {
	port := os.Getenv("PORT2")
	return ":" + port
}

func main() {
	handler := api.NewHandler()
	router := api.NewRouter(handler)
	server := http.Server{
		Addr:    getPort(),
		Handler: router.GetMux(),
	}
	log.Println("server started on port ", getPort())
	log.Println(server.ListenAndServe())
}
