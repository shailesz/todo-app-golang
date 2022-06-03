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

	// file server
	fs := http.FileServer(http.Dir("./src/client/assets"))

	mux.Handle("/src/client/assets/", http.StripPrefix("/src/client/assets/", fs))

	// start server
	log.Println("Server starting in port: " + port)
	http.ListenAndServe(":"+port, mux)
}
