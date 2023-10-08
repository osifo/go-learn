package main

import (
	"fmt"
	"net/http"
)

func getHello(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("API Server is running."))
}

func index(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("This is the Go API Server."))
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", getHello)
	server.HandleFunc("/hello", index)

	err := http.ListenAndServe(":4000", server)
	if err != nil {
		fmt.Println("Server could not be opened on port 4000")
	}
}
