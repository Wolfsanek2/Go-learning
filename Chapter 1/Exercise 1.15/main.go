// function design with pointers
package main

import "fmt"

func add5Value(count int) {
	count += 5
	fmt.Printf("add5Value: %#v\n", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Printf("add5Point: %#v\n", *count)
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:", count)
	add5Point(&count)
	fmt.Println("add5Point post:", count)
}
