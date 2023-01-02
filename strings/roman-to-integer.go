package main

import "fmt"

func main() {
	s := "MCMXCIV"

	fmt.Println(RomantToIntegers(s))
}

func RomantToIntegers(s string) int {
	Roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && Roman[string(s[i])] < Roman[string(s[i+1])] {
			res -= Roman[string(s[i])]
		} else {
			res += Roman[string(s[i])]
		}
	}

	return res
}
