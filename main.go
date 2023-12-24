package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	message, ok := data["message"]
	if !ok {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Println(message) 

	response := map[string]string{
		"status":  "success",
		"message": "Data successfully received",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	http.ListenAndServe(":8080", nil)
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	http.ListenAndServe(":8080", nil)
}
