package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{9, 1, 1, 4, 2, 5, 9, 7, 5, 2, 9, 9, 9, 10, 15, 12}

	fmt.Println(findLargest(arr))
}

func findLargest(arr []int) int {
	sort.Ints(arr)
	c := false
	first := 0
	for i := len(arr) - 2; i >= 0; i-- {
		if !c && arr[i] == arr[i+1] {
			c = true
			first = arr[i]
		}
		if c && arr[i] == arr[i+1] && arr[i] != first {
			return arr[i]
		}
	}
	return -1
}
