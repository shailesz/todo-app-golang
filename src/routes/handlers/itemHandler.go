package handlers

import (
	"fmt"
	"net/http"
	"todo/src/controllers"
	"todo/src/models"
)

// connect mongodb
var MongoClient = controllers.GetClient()

// get collection handle
var TodosCollection = controllers.GetConnectionHandle(MongoClient, "todos")

// AddItem handles route to add an item.
func AddItem(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		description := r.FormValue("description")

		x := models.NewItem(description)
		controllers.InsertDocument(TodosCollection, x)

		w.Write(x.ToJson(w))

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
