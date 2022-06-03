package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// introduce templates
var tpl = template.Must(template.ParseFiles("./src/client/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	// handle port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// multiplexer
	mux := http.NewServeMux()

	// handle routes
	mux.HandleFunc("/", indexHandler)

	// start server
	log.Println("Server starting in port: " + port)
	http.ListenAndServe(":"+port, mux)
}
