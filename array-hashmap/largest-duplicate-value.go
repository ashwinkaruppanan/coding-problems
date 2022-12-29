package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{9, 4, 2, 5, 9, 7, 5, 2, 10, 15, 12}

	sort.Ints(arr)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] == arr[i+1] {
			fmt.Println(arr[i])
			break
		}
	}
}
