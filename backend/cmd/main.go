package main

import (
	"fmt"
	"net/http"
	"os"

	"groupie-trackers/backend/handlers"
)

func main() {
	port := os.Args[1]
	// Set up the HTTP server routes
	http.HandleFunc("/static/", handlers.ServeStyle)
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/Artists", handlers.Artists)

	// Start the HTTP server
	fmt.Println("the server is running on localhost port 9999")
	fmt.Println("http://localhost:"+port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
