package main

import "fmt"

func rotate(nums []int, k int) {
	if len(nums) == 1 || k == 0 {
		fmt.Println(nums)
		return
	}

	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		next := nums[0]
		nums[0] = last
		for j := 1; j < len(nums); j++ {
			temp := nums[j]
			nums[j] = next
			next = temp
		}
		fmt.Println(nums)
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// [1,2,3,4,5,6,7]
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

// This is a working solution but with O(n × k) -> so is not efficient specially for larger inputs with k > 50K
