// Creating slices from a slice
package main

import "fmt"

func message() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := fmt.Sprintln("First:", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("Last:", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	m += fmt.Sprintln("First 5:", s[:5])
	m += fmt.Sprintln("Last 4:", s[5:])
	m += fmt.Sprintln("Middle 5:", s[2:7])
	return m
}

func main() {
	fmt.Print(message())

	// При создании нового слайса с помощью [] внутренний массив не копируется
	s1 := []int{0, 1, 2, 3, 4}
	s2 := s1[1:4]
	fmt.Println(s1, s2)
	s1[0] = 9
	s1[1] = 8
	s1[2] = 7
	s1[3] = 6
	s1[4] = 5
	fmt.Println(s1, s2)
}
