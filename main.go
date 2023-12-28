package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handlePost)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if requestBody.Message == "" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Println("Received message:", requestBody.Message)

	response := ResponseBody{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
