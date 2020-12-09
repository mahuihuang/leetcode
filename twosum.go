package main

func main() {
	nums := []int{2,7,11,15}
	_ = twoSum(nums, 9)

}

func twoSum(nums []int, target int) []int {
	var arr = make([]int,2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] + nums[j] == target {
				arr[0] = i
				arr[1] = j
				return arr
			}
		}
	}

	return arr
}