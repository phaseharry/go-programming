package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	// getting the HTML of google.com
	res, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	// closing the Body connection once we're done with it as it is streamed into server and connection is still open
	// even if your function is done with it.
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// converting the byte data we get from the GET call into a string
	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Println(data)
}
