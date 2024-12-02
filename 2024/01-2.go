package main

import (
	"fmt"
	"slices"
	"strconv"
)

func Day1_2() {
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
		in_right := 0
		for j := range len(right) {
			if left[i] == right[j] {
				in_right += 1
			}
		}
		result += left[i] * in_right
	}
	fmt.Println(result)
}
