package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("This is the Go API Server."))
}

func main() {
	server := http.NewServeMux()

	website := http.FileServer(http.Dir("./public"))

	server.Handle("/", website)
	server.HandleFunc("/hello", index)

	err := http.ListenAndServe(":4000", server)
	if err != nil {
		fmt.Println("Server could not be opened on port 4000")
	}
}
