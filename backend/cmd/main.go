package main

import (
	"fmt"
	"net/http"

	"groupie-trackers/backend/handlers"
)

func main() {
	// Set up the HTTP server routes
	http.HandleFunc("/static/", handlers.ServeStyle)
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/Artists", handlers.Artists)

	// Start the HTTP server
	fmt.Println("the server is running on localhost port 7777")
	fmt.Println("http://localhost:7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println(err)
	}
}
