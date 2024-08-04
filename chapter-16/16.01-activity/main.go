package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	Counter int
	Content string
	Heading string
}

func (h *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Counter++
	msg := fmt.Sprintf("<h1>%s</h1><p>%s</p><div>%d views</div>", h.Heading, h.Content, h.Counter)
	w.Write([]byte(msg))
}

func main() {
	homePage := PageWithCounter{
		Counter: 0,
		Heading: "Home Page",
		Content: "This is the homepage.",
	}
	chapter1 := PageWithCounter{
		Counter: 0,
		Heading: "Chapter 1",
		Content: "This is the chapter 1.",
	}
	chapter2 := PageWithCounter{
		Counter: 0,
		Heading: "Chapter 2",
		Content: "This is the chapter 2.",
	}
	http.Handle("/", &homePage)
	http.Handle("/chapter1", &chapter1)
	http.Handle("/chapter2", &chapter2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
