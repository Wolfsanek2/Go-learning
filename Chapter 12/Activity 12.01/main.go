// Formatting a date according to user requirements
package main

import (
	"fmt"
	"strconv"
	"time"
)

func currentDate() string {
	now := time.Now()
	hours := strconv.Itoa(now.Hour())
	minutes := strconv.Itoa(now.Minute())
	seconds := strconv.Itoa(now.Second())
	day := strconv.Itoa(now.Day())
	month := strconv.Itoa(int(now.Month()))
	year := strconv.Itoa(now.Year())
	result := fmt.Sprintf("%s:%s:%s %s/%s/%s", hours, minutes, seconds, day, month, year)
	return result
}

func main() {
	fmt.Println(currentDate())
}
