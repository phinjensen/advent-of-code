package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3_1() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
	input := string(bytes)

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	result := 0
	for _, match := range matches {
		l, _ := strconv.Atoi(match[1])
		r, _ := strconv.Atoi(match[2])
		result += l * r
	}

	fmt.Println(result)
}

func Day3_2() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
	input := string(bytes)

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	result := 0
	do := true
	for _, match := range matches {
		fmt.Println(match)
		if do && strings.HasPrefix(match[0], "mul") {
			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])
			result += l * r
		} else if strings.HasPrefix(match[0], "don't") {
			do = false
		} else if strings.HasPrefix(match[0], "do") {
			do = true
		}
	}

	fmt.Println(result)
}
