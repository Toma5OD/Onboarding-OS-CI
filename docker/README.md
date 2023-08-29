# HTTPS Echo Server in Go with Docker

## Description

This Go program creates an HTTP server that listens on port 8080. The server responds by echoing the received request body back to the client. The program is also containerized using Docker for easy deployment.

## How It Works

- Uses Go's `net/http` package to create an HTTP server.
- Listens on port 8080 and handles requests at the `/echo` endpoint.
- Uses `io.Copy` to echo the request body back to the client.
- Containerized with Docker to run inside a Podman container.

## Features

- Echos back the request body received at `/echo`.
- Serves a static HTML file (`index.html`) for user interaction.
  
## How to Use

### Prerequisites

Make sure you have Go, Docker, and Podman installed on your machine.

### Prepare Go Files

Place your Go source files (`http-echo-server.go`) and any HTML files (`index.html`) in the `docker` directory.

```bash
cd docker/
```

### Build the Docker Image

Navigate to the `docker` directory and run the following command:

```bash
podman build -t https-echo-server .
```

### Run the Container

After the build is successful, run the following command:

```bash
podman run -p 8080:8080 https-echo-server
```

### Test the Server

Navigate to `http://localhost:8080` in your web browser to interact with the server.

### Exec into the Running Container

Find the container ID using `podman ps` and then exec into it using:

```bash
podman exec -it <container_id> /bin/sh
```

### Check Logs

You can check the logs of the running container using:

```bash
podman logs <container_id>
```

## Notes

- If you're running the server locally, replace `localhost` with the appropriate IP address if needed.

Feel free to expand upon this README for more specific requirements or features related to your use case.
