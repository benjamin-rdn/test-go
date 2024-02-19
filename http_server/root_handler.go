package main

import (
	"html/template"
	"net/http"
)

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	// Get the URL parameter "name"
	urlParam := request.URL.Query().Get("name")
	// Parse the HTML file
	tpl := template.Must(template.ParseFiles("index.html"))
	// Define the content to be displayed in the template
	content := map[string]string{"content": "Hello " + urlParam}
	// Execute the template
	tpl.Execute(writer, content)
}
