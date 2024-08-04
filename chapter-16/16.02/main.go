package main

import (
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	w.Write([]byte(msg))
}

func main() {
	// will handle all paths with /chapter1
	http.HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request) {
		msg := "Chapter 1"
		w.Write([]byte(msg))
	})
	// will handle root paths and any other path besides "/chapter1"
	http.Handle("/", hello{})

	/*
		Difference between .Handle vs .HandleFunc.
		- HandleFunc takes in a function with signature (w http.ResponseWriter, r *http.Request)
		- Handle takes in a struct that implements the http.Handler interface. (i.e. implement the ServeHTTP(w http.ResponseWriter, r *http.Request) method)
		Both are used to determine what function is called for a specified path. The choice in which to pick is usually based on personal preference, but for more
		complex routing needs, .Handle should be used to do it. Different routes may appear better organized if they're handled by handlers that belong to different packages.
		If routes have little functionality, then a simple function (.HandleFunc) could be used.
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}
