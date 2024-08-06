package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
a struct that a json response will me marshalled into so we can directly use it in our code.
*/
type messageData struct {
	Message string `json:"message"`
}

func getDataAndReturnResponse() messageData {
	/*
		assuming r.Body is a JSON payload like the following:
		{
		"message": "Hello World"
		}
		The key "message"'s value will be marshalled into our messageData struct's "Message" field.
		We made "Message" titled cased so the key value can be exported and used.
	*/
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	// reads the data sent and converts it to a slice of bytes
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	message := messageData{}
	// passing in the raw byte data and message struct so the raw byte data can be marshalled into the struct
	// and can be used within our program
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Println(data.Message)
}
