// looping over map data using range
package main

import "fmt"

func main() {
	data := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}
	maxKey := ""
	maxValue := 0
	for key, value := range data {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}
	fmt.Println(maxKey, maxValue)
}
