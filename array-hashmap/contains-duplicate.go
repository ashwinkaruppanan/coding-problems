package main

import "fmt"

func main() {
	arr := []int{9, 4, 2, 5, 9, 7, 5, 2, 10, 15, 12}

	fmt.Println(containsDuplicate(arr))
}

func containsDuplicate(nums []int) bool {
	record := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		_, isExist := record[nums[i]]
		if isExist {
			return true
		} else {
			record[nums[i]] = true
		}
	}
	fmt.Println(record)
	return false
}
