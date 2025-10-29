package main

import (
	"fmt"
)

func rotateS3(nums []int, k int) {

	if k == 0 || len(nums) == 1 || k == len(nums) {
		fmt.Println(nums)
		return
	}

	rotationCount := k % len(nums)

	last := nums[len(nums)-1]
	newSlice := make([]int, k)
	newSlice[k-1] = last
	c := k
	for i := 1; i < rotationCount; i++ {
		last = nums[len(nums)-c]
		c = c - 1
		newSlice[c-1] = last
	}
	restOfSlice := nums[0 : len(nums)-k]
	fmt.Println(newSlice)
	fmt.Println(restOfSlice)
}

func main() {
	rotateS3([]int{-1, -100, 3, 99}, 2)
}

// Examples

// [1,2,3,4,5,6,7]
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
