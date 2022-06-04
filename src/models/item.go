package models

import (
	"encoding/json"
	"net/http"
	"time"
)

// Item is a todo item.
type Item struct {
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
	Timestamp   int64  `json:"timestamp"`
}

// NewItem is constructor for Item given a description.
func NewItem(description string) Item {
	x := Item{Description: description, IsComplete: false, Timestamp: time.Now().UnixNano()}

	return x
}

// ToJson marshals Item to Json.
func (i *Item) ToJson(w http.ResponseWriter) []byte {
	itemJson, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return itemJson
}
