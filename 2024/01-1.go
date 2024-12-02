package main

import (
	"fmt"
	"slices"
	"strconv"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	var a, b string
	var left, right []int
	for {
		_, err := fmt.Scanln(&a, &b)
		if err != nil {
			break
		}
		a_int, _ := strconv.Atoi(a)
		b_int, _ := strconv.Atoi(b)
		left = append(left, a_int)
		right = append(right, b_int)
	}
	slices.Sort(left)
	slices.Sort(right)
	result := 0
	for i := range len(left) {
		result += abs(left[i] - right[i])
	}
	fmt.Println(result)
}
