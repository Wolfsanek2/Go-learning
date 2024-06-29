// Creating a locale checker
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

var locales = []locale{
	locale{
		language: "en",
		region:   "US",
	},
	{
		"en",
		"CN",
	},
	{
		"fr",
		"CN",
	},
	{
		"fr",
		"Fr",
	},
	{
		"ru",
		"RU",
	},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter locale")
		os.Exit(1)
	}
	locale_str := os.Args[1]
	locale_arr := strings.Split(locale_str, "_")
	if len(locale_arr) != 2 {
		fmt.Println("Error")
		os.Exit(1)
	}
	loc := locale{
		locale_arr[0],
		locale_arr[1],
	}
	var exists bool
	for _, cur := range locales {
		if cur == loc {
			exists = true
			break
		}
	}
	if exists {
		fmt.Println("Locale supported")
	} else {
		fmt.Println("Locale not supported")
	}
}
