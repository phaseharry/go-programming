package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Hello(writer http.ResponseWriter, request *http.Request) {
	/*
		For routes that maps to this function, it parses it query string parameters.
		ex) localhost:8080?name=Harry
	*/
	query := request.URL.Query()
	/*
		checks if "name" is part of the query string. If it is part of the query string, then "ok"
		will be true. In the below case, if ok is false then we write a 400 bad request back to client.
	*/
	name, ok := query["name"]
	if !ok {
		writer.WriteHeader(400)
		writer.Write([]byte("Missing name"))
		return
	}
	// "name" is slice of strings even if it's just one value so we need to Join it into one string and output that.
	writer.Write([]byte(fmt.Sprintf("Hello %s", strings.Join(name, ","))))
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
