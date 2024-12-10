package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func Day9_1() {
	input, _ := io.ReadAll(os.Stdin)
	disk_map := string(input)
	files := make([]int, 0)
	space := make([]int, 0)
	for i, c := range disk_map {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			files = append(files, n)
		} else {
			space = append(space, n)
		}
	}
	slices.Reverse(space)
	i := 0
	front := 0
	back := len(files) - 1
	in_space := false
	checksum := 0
	for front < back {
		if in_space {
			// didn't even account for the spaces yet :(
			if files[back] == 0 {
				back -= 1
				in_space = false
			} else {
				checksum += i * back
				files[back] -= 1
				i++
			}
		} else {
			if files[front] == 0 {
				front += 1
				in_space = true
			} else {
				checksum += i * front
				files[front] -= 1
				i++
			}
		}
	}
	fmt.Println(checksum)
}
