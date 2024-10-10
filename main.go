package main

import (
	"app/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/task/add", handlers.TaskHandler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
