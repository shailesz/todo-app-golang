package handlers

import (
	"html/template"
	"net/http"
	"todo/src/controllers"
)

// introduce templates
var tpl = template.Must(template.ParseFiles("./src/client/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, controllers.GetAllDocuments(TodosCollection))
}
