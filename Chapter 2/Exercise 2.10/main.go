// looping over a map
package main

import "fmt"

func main() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}
	for key, value := range config {
		fmt.Println(key, "=", value)
	}
	for key, _ := range config {
		fmt.Println(key)
	}
	for _, value := range config {
		fmt.Println(value)
	}
}
