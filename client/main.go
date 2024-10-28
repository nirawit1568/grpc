package main

import (
	"client/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Run server for client connection
	fmt.Println("Server is running on http://localhost:8080")
	http.HandleFunc("/hello", handlers.Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
