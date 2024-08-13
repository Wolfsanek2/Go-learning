// using the go vet tool
package main

import "fmt"

func main() {
	helloString := "Hello"
	packtString := "Pactk"
	jointString := fmt.Sprintf("%s %s", helloString, packtString)
	fmt.Println(jointString)
}
