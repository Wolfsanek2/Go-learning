// Removing an element from a slice
package main

import "fmt"

var data []string = []string{
	"Good",
	"Good",
	"Bad",
	"Good",
	"Good",
}

func remove() {
	var index int
	for i, value := range data {
		if value == "Bad" {
			index = i
		}
	}
	data[index] = data[len(data)-1]
	data = data[:len(data)-1]
}

func main() {
	remove()
	fmt.Println(data)
}
