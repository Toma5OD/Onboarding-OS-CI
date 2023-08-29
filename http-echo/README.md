# Go HTTP Echo Server

## Description

This Go program creates a simple HTTP server that echoes back the content it receives in a POST request.

## How It Works

- Listens for HTTP POST requests at the `/echo` route using Go's `net/http` package.
- Reads the body of the incoming POST request using the `io` package.
- Writes the read content back to the HTTP response, essentially echoing the request.

## How to Use

1. Ensure you have Go installed on your machine.
2. Navigate to the directory containing the Go program.
3. Run the program with the command `go run main.go`.
4. Server starts at `http://localhost:8080`. Test it by sending a POST request to `http://localhost:8080/echo` with some body content.
5. The server will echo back the content it received in the POST request.

## Code Overview

- `http.HandleFunc("/echo", echoHandler)`: Registers the `echoHandler` function to handle requests to the `/echo` route.
- `http.ListenAndServe(":8080", nil)`: Starts the HTTP server on port 8080.
- `io.Copy(w, r.Body)`: Reads the request body and writes it directly back to the response.
