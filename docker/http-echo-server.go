package main

import (
	"io"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, r.Body)
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
