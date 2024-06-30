// Type checker
package main

import "fmt"

func typeChecker(v interface{}) string {
	switch v.(type) {
	case int, int32, int64:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println("1 is", typeChecker(1))
	fmt.Println("3.14 is", typeChecker(3.14))
	fmt.Println("hello is", typeChecker("hello"))
	fmt.Println("true is", typeChecker(true))
	fmt.Println("{} is", typeChecker(struct{}{}))
}
