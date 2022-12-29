package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "anagrams"
	t := "nagaramw"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	S := strings.Split(s, "")
	T := strings.Split(t, "")

	sort.Strings(S)
	sort.Strings(T)
	for i := 0; i < len(S); i++ {
		if S[i] != T[i] {
			return false
		}
	}
	return true
}
