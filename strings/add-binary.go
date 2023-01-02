package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "0"
	b := "1"

	fmt.Println(AddBinary(a, b))
}

func AddBinary(a string, b string) string {

	if len(a) < len(b) {
		a, b = b, a
	}
	if len(a) != len(b) {
		for i := len(b); i < len(a); i++ {
			b = "0" + b
		}
	}

	a = reverse(a)
	b = reverse(b)

	output := ""
	carry := 0

	for i := 0; i < len(a); i++ {
		dig1 := int(a[i]) - '0'
		dig2 := int(b[i]) - '0'

		total := dig1 + dig2 + carry
		output = strconv.Itoa(total%2) + output
		carry = total / 2

	}

	if carry != 0 {
		output = "1" + output
	}
	return output
}

func reverse(a string) string {
	output := ""
	for i := len(a) - 1; i >= 0; i-- {
		output += string(a[i])
	}
	return output
}
