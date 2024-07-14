// Enforcing a specific format of date and time
package main

import (
	"fmt"
	"strconv"
	"time"
)

func date() string {
	date := time.Date(2022, 2, 24, 10, 50, 30, 0, time.Local)
	hours := strconv.Itoa(date.Hour())
	minutes := strconv.Itoa(date.Minute())
	seconds := strconv.Itoa(date.Second())
	day := strconv.Itoa(date.Day())
	month := strconv.Itoa(int(date.Month()))
	year := strconv.Itoa(date.Year())
	result := fmt.Sprintf("%s:%s:%s %s/%s/%s", hours, minutes, seconds, day, month, year)
	return result
}

func main() {
	fmt.Println(date())
}
