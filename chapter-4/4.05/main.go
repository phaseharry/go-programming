package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language string
	region   string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No locale passed in!")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	inputLocale := locale{
		language: localeParts[0],
		region:   localeParts[1],
	}

	availableLocales := getLocaleCollection()
	for i := 0; i < len(availableLocales); i++ {
		if availableLocales[i] == inputLocale {
			fmt.Println("Locale passed in is supported!")
			return
		}
	}
	fmt.Println("Locale passed in is not supported!")
}

func getLocaleCollection() []locale {
	return []locale{
		{
			language: "en",
			region:   "US",
		},
		{
			language: "en",
			region:   "CN",
		},
		{
			language: "fr",
			region:   "CN",
		},
		{
			language: "fr",
			region:   "FR",
		},
		{
			language: "ru",
			region:   "RU",
		},
	}
}
