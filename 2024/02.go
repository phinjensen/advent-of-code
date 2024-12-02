package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day2_1() {
	safe := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		values_s := strings.Split(scanner.Text(), " ")
		values := make([]int, len(values_s))
		for i, e := range values_s {
			values[i], _ = strconv.Atoi(e)
		}
		ascending := false
		line_safe := true
		for i := range len(values) - 1 {
			if i == 0 {
				if values[i] < values[i+1] {
					ascending = true
				} else if values[i] > values[i+1] {
					ascending = false
				} else {
					line_safe = false
					break
				}
			}

			diff := values[i] - values[i+1]
			if ascending {
				diff = -diff
				if diff < 1 || diff > 3 {
					line_safe = false
					break
				}
			} else {
				if diff < 1 || diff > 3 {
					line_safe = false
					break
				}
			}
		}
		if line_safe {
			safe += 1
		}
	}
	fmt.Println(safe)
}

func Day2_2() {
	safe := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		values_s := strings.Split(scanner.Text(), " ")
		values := make([]int, len(values_s))
		for i, e := range values_s {
			values[i], _ = strconv.Atoi(e)
		}
		ascending := false
		fmt.Println(values)
		for remove := range len(values) {
			line_safe := true
			temp_values := slices.Concat(values[:remove], values[remove+1:])
			fmt.Println("\t", temp_values)
			for i := range len(temp_values) - 1 {
				if i == 0 {
					if temp_values[i] < temp_values[i+1] {
						ascending = true
					} else if temp_values[i] > temp_values[i+1] {
						ascending = false
					} else {
						line_safe = false
						break
					}
				}

				diff := temp_values[i] - temp_values[i+1]
				if ascending {
					diff = -diff
				}
				if diff < 1 || diff > 3 {
					line_safe = false
					break
				}
			}
			if line_safe {
				safe += 1
				fmt.Println("")
				break
			}
		}
	}
	fmt.Println(safe)
}
