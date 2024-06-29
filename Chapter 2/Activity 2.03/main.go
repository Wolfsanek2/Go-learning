// bubble sort
package main

import "fmt"

func main() {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	fmt.Println(nums)
}
