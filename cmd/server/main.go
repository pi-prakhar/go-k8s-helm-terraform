package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/pi-prakhar/go-k8s-helm-terraform/internal"
)

func getPort() string {
	port := os.Getenv("PORT")
	return ":" + port
}

func main() {
	handler := api.NewHandler()
	router := api.NewRouter(handler)
	server := http.Server{
		Addr:    getPort(),
		Handler: router.GetMux(),
	}

	log.Println(server.ListenAndServe())
}
