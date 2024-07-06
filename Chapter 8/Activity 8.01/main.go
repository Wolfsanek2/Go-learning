// a minimum value
package main

import (
	"errors"
	"fmt"
)

func findMinGeneric[Num int | float64](nums []Num) (Num, error) {
	if len(nums) == 0 {
		return 0, errors.New("empty slice")
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min, nil
}

func main() {
	fmt.Println(findMinGeneric([]int{}))
	fmt.Println(findMinGeneric([]int{1, 32, 5, 8, 10, 11}))
	fmt.Println(findMinGeneric([]float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1}))
}
