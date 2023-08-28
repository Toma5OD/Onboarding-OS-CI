package main

import (
	"io"
	"net/http"
)

// echoHandler handles the /echo route
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Use io.Copy to read the request body and write it back to the response
	// This essentially "echoes" the request back to the client
	io.Copy(w, r.Body)
}

func main() {
	// Register the echoHandler function for the "/echo" route
	http.HandleFunc("/echo", echoHandler)
	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
