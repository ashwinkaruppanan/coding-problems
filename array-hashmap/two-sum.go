package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(arr, target))
}

func twoSum(nums []int, target int) []int {
	record := map[int]int{}

	for i := 0; i < len(nums); i++ {
		index, isExist := record[target-nums[i]]
		if isExist {
			return []int{index, i}
		} else {
			record[nums[i]] = i
		}
	}
	return nil
}
