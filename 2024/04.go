package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func CheckDirection(r int, c int, input []string, direction [4][2]int) bool {
	word := ""
	for _, d := range direction {
		row := r + d[0]
		col := c + d[1]
		if row < 0 || col < 0 || row >= len(input) || col >= len(input[row]) {
			return false
		}
		word += string(input[row][col])
	}
	if word == "XMAS" || word == "SAMX" {
		return true
	}
	return false
}

func Day4_1() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
	lines := strings.Split(string(bytes), "\n")
	total := 0
	directions := [4][4][2]int{
		{{0, 0}, {-1, -1}, {-2, -2}, {-3, -3}}, // nw
		{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}},    // n
		{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}},    // ne
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},       // e
		//{{0, 0}, {1, 1}, {2, 2}, {3, 3}},       // se
		//{{0, 0}, {1, 0}, {2, 0}, {3, 0}},       // s
		//{{0, 0}, {1, -1}, {2, -2}, {3, -3}},    // sw
		//{{0, 0}, {0, -1}, {0, -2}, {0, -3}},    // w
	}
	for r := range lines {
		for c := range lines {
			for _, direction := range directions {
				if CheckDirection(r, c, lines, direction) {
					total += 1
				}
			}
		}
	}
	fmt.Println(total)
}

func CheckX(r int, c int, input []string) bool {
	directions := [2][3][2]int{
		{{-1, -1}, {0, 0}, {1, 1}},
		{{1, -1}, {0, 0}, {-1, 1}},
	}
	for _, direction := range directions {
		word := ""
		for _, d := range direction {
			row := r + d[0]
			col := c + d[1]
			if row < 0 || col < 0 || row >= len(input) || col >= len(input[row]) {
				return false
			}
			word += string(input[row][col])
		}
		if !(word == "MAS" || word == "SAM") {
			return false
		}
	}
	fmt.Println(r, c)
	return true
}

func Day4_2() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
	lines := strings.Split(string(bytes), "\n")
	total := 0
	for r := range lines {
		for c := range lines {
			if CheckX(r, c, lines) {
				total += 1
			}
		}
	}
	fmt.Println(total)
}
