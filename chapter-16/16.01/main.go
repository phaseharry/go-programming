package main

import (
	"log"
	"net/http"
)

type helloWorldServer struct{}

func (h helloWorldServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	w.Write([]byte(msg))
}

func main() {
	/*
		helloWorldServer struct is implementing the httpHandler interface. It needs to have a ServerHTTP method
		with a function signature of (http.ResponseWriter, http.Request) to be accepted by the interface.

		starts an HTTP server on port 8080. In this case, all routes to this server will response with "<h1>Hello World</h1>"

		http.ListenAndServe method might return an error. If that happens, we want the main function's execution to stop.
		ListenAndServe won't return anything unless it errored out. If there's no error, the server keeps running and nothing is returned.
		In the below case, we call log.Fatal to log the error and exit the program.

	*/
	log.Fatal(http.ListenAndServe(":8080", helloWorldServer{}))
}
