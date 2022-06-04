package main

import (
	"log"
	"net/http"
	"os"
	"todo/src/controllers"
	"todo/src/routes/handlers"
	"todo/src/services"
)

func main() {
	// handle port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := services.InitApp()

	// someItem := models.Item{Description: string(time.Now().UnixNano()), IsComplete: false, Timestamp: time.Now().UnixNano()}

	// insert a document
	// controllers.InsertDocument(todosCollection, someItem)

	// get all documents
	controllers.GetAllDocuments(handlers.TodosCollection)

	// start server
	log.Println("Server starting in port: " + port)
	http.ListenAndServe(":"+port, mux)
}
