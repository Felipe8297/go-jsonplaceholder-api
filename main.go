package main

import (
	"api/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.GetPostsHandler)
	fmt.Println("Server running at http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}
