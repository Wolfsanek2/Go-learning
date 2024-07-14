// Calculating the future date and time
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.ANSIC))
	duration := 6*time.Hour + 6*time.Minute + 6*time.Second
	future := now.Add(duration)
	fmt.Println(future.Format(time.ANSIC))
}
