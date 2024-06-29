// expressionless switch statements
package main

import (
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Born some other day")
	}

	var a int = 5
	b := 8
	switch {
	case a > 0 && a < 5:
		fmt.Println(a)
	case a < b:
		fmt.Println("a < b")
	default:
		fmt.Println("b")
	}
}
