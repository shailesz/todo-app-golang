package main

import (
	"log"
	"net/http"
	"os"
	"todo/src/services"
)

func main() {
	// handle port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// init app
	mux := services.InitApp()

	// start server
	log.Println("Server starting in port: " + port)
	http.ListenAndServe(":"+port, mux)
}
