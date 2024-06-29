// declaring multiple variables from a function
package main

import (
	"fmt"
	"time"
)

func main() {
	Debug, LogLevel, startUpTime := getConfig()

	fmt.Println(Debug, LogLevel, startUpTime)
}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
