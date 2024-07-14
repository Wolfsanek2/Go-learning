// Printing the local time in different time zones
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timeZone1, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	timeZone2, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}
	NyTime := now.In(timeZone1)
	LaTime := now.In(timeZone2)
	fmt.Println(now.Format(time.ANSIC))
	fmt.Println(NyTime.Format(time.ANSIC))
	fmt.Println(LaTime.Format(time.ANSIC))
}
