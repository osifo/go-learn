package main

import (
	"fmt"
	"html/template"
	"net/http"

	"osifo.dev/go-learn/web/data"
)

func index(writer http.ResponseWriter, reqeust *http.Request) {
	writer.Write([]byte("This is the Go API Server."))
}

func handleTemplate(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("./templates/index.tmpl"))
	fmt.Printf("%v", data.GetLatestExhibition())
	templateError := template.Execute(writer, data.GetAllExhibitions())

	if templateError != nil {
		errMessage := fmt.Sprintf("Could not pass template. Error:\n%v\n", templateError)
		print(errMessage)

		// writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(errMessage))
	}
}

func main() {
	server := http.NewServeMux()

	// website := http.FileServer(http.Dir("./public"))
	// server.Handle("/", website)

	server.HandleFunc("/", handleTemplate)
	server.HandleFunc("/hello", index)

	err := http.ListenAndServe(":4000", server)
	if err != nil {
		fmt.Println("Server could not be opened on port 4000")
	}
}
