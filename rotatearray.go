package main

import "fmt"

func main() {
	array := []int{7,1,2,3,4,5,6}
	k := 3

	rotate(array, k)
}

func rotate(nums []int, k int)  {
	for i := 0; i < k; i++ {
		max := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = max
	}
}