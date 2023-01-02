package main

import (
	"fmt"
)

func main() {
	inp := 5

	fmt.Println(generate(inp))
}

func generate(numRows int) [][]int {
	output := [][]int{{1}}

	for i := 1; i < numRows; i++ {
		temp := []int{0}
		temp = append(temp, output[len(output)-1]...)
		temp = append(temp, []int{0}...)
		row := []int{}
		for j := 0; j < len(output[len(output)-1])+1; j++ {
			row = append(row, temp[j]+temp[j+1])
		}
		output = append(output, row)

	}

	return output
}
