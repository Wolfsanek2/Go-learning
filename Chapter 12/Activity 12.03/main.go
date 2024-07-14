// Measuring elapsed time
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	duration := 2 * time.Second
	time.Sleep(duration)
	end := time.Now()
	dif := end.Sub(start)
	fmt.Println(dif)
}
