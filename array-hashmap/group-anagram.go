package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {

	maps := make(map[string][]string)

	for _, val := range strs {
		byteString := []byte(val)
		sort.Slice(byteString, func(i, j int) bool { return byteString[i] < byteString[j] })
		sortedString := string(byteString)

		maps[sortedString] = append(maps[sortedString], val)
	}

	ans := [][]string{}

	for _, val := range maps {
		ans = append(ans, val)
	}

	return ans
}
