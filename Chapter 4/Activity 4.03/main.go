// Slicing the week
package main

import "fmt"

func main() {
	week := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	week = append(week[6:], week[:6]...)
	fmt.Println(week)
}
