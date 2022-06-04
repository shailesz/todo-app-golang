package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"todo/src/controllers"
	"todo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// connect mongodb
var MongoClient = controllers.GetClient()

// get collection handle
var TodosCollection = controllers.GetConnectionHandle(MongoClient, "todos")

// AddItem handles route to add an item.
func AddItem(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		headerContentTtype := r.Header.Get("Content-Type")

		if headerContentTtype != "application/json" {
			sendResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}

		var item models.Item

		decodeRequest(r, &item)

		x := models.NewItem(item.Description)

		controllers.InsertDocument(TodosCollection, x)

		w.Write(x.ToJson(w))

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// DeleteItem handles route to delete an item.
func DeleteItem(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		headerContentTtype := r.Header.Get("Content-Type")

		if headerContentTtype != "application/json" {
			sendResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}

		var item models.Item

		decodeRequest(r, &item)

		controllers.DeleteDocument(TodosCollection, item.Id)

		sendResponse(w, "Success!", http.StatusOK)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// MarkComplete marks passed items as complete.
func MarkComplete(w http.ResponseWriter, r *http.Request) {

	type Items struct {
		MarkComplete []string `json:"markComplete"`
	}

	switch r.Method {
	case "POST":

		headerContentTtype := r.Header.Get("Content-Type")

		if headerContentTtype != "application/json" {
			sendResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}

		var data Items
		var objectIds []primitive.ObjectID

		decodeRequest(r, &data)

		for _, x := range data.MarkComplete {
			o, err := primitive.ObjectIDFromHex(x)
			if err != nil {
				panic(err.Error())
			}
			objectIds = append(objectIds, o)
		}

		filter, update := bson.M{"_id": bson.M{"$in": objectIds}}, bson.D{{"$set", bson.D{{"iscomplete", true}}}}

		_, err := TodosCollection.UpdateMany(
			context.TODO(),
			filter,
			update,
		)
		if err != nil {
			log.Fatal(err)
		}
		sendResponse(w, "Success!", http.StatusOK)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// UnmarkComplete marks passed items as incomplete.
func UnmarkComplete(w http.ResponseWriter, r *http.Request) {

	type Items struct {
		UnmarkComplete []string `json:"unmarkComplete"`
	}

	switch r.Method {
	case "POST":

		headerContentTtype := r.Header.Get("Content-Type")

		if headerContentTtype != "application/json" {
			sendResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}

		var data Items
		var objectIds []primitive.ObjectID

		decodeRequest(r, &data)

		for _, x := range data.UnmarkComplete {
			o, err := primitive.ObjectIDFromHex(x)
			if err != nil {
				panic(err.Error())
			}
			objectIds = append(objectIds, o)
		}

		filter, update := bson.M{"_id": bson.M{"$in": objectIds}}, bson.D{{"$set", bson.D{{"iscomplete", false}}}}

		_, err := TodosCollection.UpdateMany(
			context.TODO(),
			filter,
			update,
		)
		if err != nil {
			log.Fatal(err)
		}
		sendResponse(w, "Success!", http.StatusOK)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// helper function to send json response back.
func sendResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

// decodeRequest decodes request body from json to interface.
func decodeRequest(r *http.Request, a interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(body, a)
	if err != nil {
		panic(err.Error())
	}
}
