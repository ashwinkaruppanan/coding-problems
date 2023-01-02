package main

import (
	"fmt"
)

func main() {
	arr := []int{-1, 1, 0, -3, 3}

	fmt.Println(productExceptSelf(arr))
}

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		output[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for j := len(nums) - 1; j >= 0; j-- {
		output[j] *= postfix
		postfix *= nums[j]
	}

	return output
}
