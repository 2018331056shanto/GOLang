package main

import (
	"fmt"
	"sort"
)

func main() {

	var nums []int

	var n int

	fmt.Scan(&n)
	for i := 0; i < n; i = i + 1 {

		var x int

		fmt.Scan(&x)

		nums = append(nums, x)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var target int

	fmt.Scan(&target)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {

		sum := nums[i] + nums[j]

		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {

			fmt.Println(nums[i], nums[j])
			break
		}
	}

}
