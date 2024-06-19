package main

import (
	"fmt"
	"os"
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
	language := os.Args[1]
	availableLocales := getLocaleCollection()
	for i := 0; i < len(availableLocales); i++ {
		if availableLocales[i].language == language {
			fmt.Println(language + " is supported in " + availableLocales[i].region)
			return
		}
	}
	fmt.Println(language + " is not supported")
}

func getLocaleCollection() []locale {
	return []locale{
		{
			language: "en_US",
			region:   "United States",
		},
		{
			language: "en_CN",
			region:   "China",
		},
		{
			language: "fr_CN",
			region:   "China",
		},
		{
			language: "fr_FR",
			region:   "France",
		},
		{
			language: "ru_RU",
			region:   "Russia",
		},
	}
}
