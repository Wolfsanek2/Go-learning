// mapping a CSV index to a column header with return values
package main

import (
	"fmt"
	"strings"
)

func csvHdrCol(hdr []string) map[int]string {
	csvIdxToCol := make(map[int]string)
	for i, v := range hdr {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee", "hours worked", "hourly rate":
			csvIdxToCol[i] = v
		}
	}
	return csvIdxToCol
}

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrCol(hdr)
	fmt.Println("Result:")
	fmt.Println(result, "\n")
	hdr2 := []string{"employee", "empid", "hours worked", "address", "manger", "hourly rate"}
	result2 := csvHdrCol(hdr2)
	fmt.Println("Result2:")
	fmt.Println(result2)
}
