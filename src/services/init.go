package services

import (
	"net/http"
	"todo/src/routes/handlers"
)

func InitApp() *http.ServeMux {

	// multiplexer
	mux := http.NewServeMux()

	// handle routes
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/add", handlers.AddItem)

	// file server
	fs := http.FileServer(http.Dir("./src/client/assets"))

	mux.Handle("/src/client/assets/", http.StripPrefix("/src/client/assets/", fs))

	return mux
}
