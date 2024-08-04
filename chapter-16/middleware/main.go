package main

import (
	"log"
	"net/http"
)

// middleware function that will be called before any other function for all paths
func Hello(next http.HandlerFunc) http.HandlerFunc {
	/*
		This is the middleware function signature. It takes in the next function to be called in the chain
		as an argument and it returns a function that dev has to write to control what common behavior it should have.
		It does that common behavior and calls next
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello there,"
		w.Write([]byte(msg)) // does it's job by writing "Hello there,"" and calling next so the next function in the chain can be executed
		next.ServeHTTP(w, r)
	}
}

func Function1(w http.ResponseWriter, r *http.Request) {
	msg := " this is function 1"
	w.Write([]byte(msg))
}

func Function2(w http.ResponseWriter, r *http.Request) {
	msg := " and now we are in function 2"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/hello1", Hello(Function1)) // utilizing the middleware for both Function1 and Function2
	http.HandleFunc("/hello2", Hello(Function2))
	log.Fatal(http.ListenAndServe(":8085", nil))

	// ex of nested middleware usage: Hello(Middleware2(Middleware3((Function2)))
}
