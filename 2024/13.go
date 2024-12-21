package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day13_1() {
	scanner := bufio.NewScanner(os.Stdin)
	var button_a, button_b XY
	var prize XY
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
		outer:
			for i := range 100 {
			inner:
				for j := range 100 {
					x := button_a.X*i + button_b.X*j
					y := button_a.Y*i + button_b.Y*j
					if x == prize.X && y == prize.Y {
						result += 3*i + j
						break outer
					} else if x > prize.X || y > prize.Y {
						break inner
					}
				}
			}
		} else if strings.HasPrefix(line, "Button ") {
			label_numbers := strings.Split(line, ": ")
			xy_string := strings.ReplaceAll(label_numbers[1], "X+", "")
			xy_string = strings.ReplaceAll(xy_string, "Y+", "")
			xy_strings := strings.Split(xy_string, ", ")
			x, _ := strconv.Atoi(xy_strings[0])
			y, _ := strconv.Atoi(xy_strings[1])
			if label_numbers[0] == "Button A" {
				button_a = XY{x, y}
			} else {
				button_b = XY{x, y}
			}
		} else if strings.HasPrefix(line, "Prize: ") {
			label_numbers := strings.Split(line, ": ")
			xy_string := strings.ReplaceAll(label_numbers[1], "X=", "")
			xy_string = strings.ReplaceAll(xy_string, "Y=", "")
			xy_strings := strings.Split(xy_string, ", ")
			x, _ := strconv.Atoi(xy_strings[0])
			y, _ := strconv.Atoi(xy_strings[1])
			prize = XY{x, y}
		}
	}
	fmt.Println(result)
}

func Day13_2() {
	scanner := bufio.NewScanner(os.Stdin)
	var button_a, button_b XY
	var prize XY
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
		outer:
			for i := 1; i < 100; i++ {
			inner:
				for j := 1; j < 100; j++ {
					x := button_a.X*j + button_b.X*i
					y := button_a.Y*j + button_b.Y*i
					if prize.X%x == 0 && prize.Y%y == 0 {
						result += 3*j*(prize.X/x) + i*(prize.Y/y)
						break outer
					} else if x > prize.X || y > prize.Y {
						break inner
					}
				}
			}
		} else if strings.HasPrefix(line, "Button ") {
			label_numbers := strings.Split(line, ": ")
			xy_string := strings.ReplaceAll(label_numbers[1], "X+", "")
			xy_string = strings.ReplaceAll(xy_string, "Y+", "")
			xy_strings := strings.Split(xy_string, ", ")
			x, _ := strconv.Atoi(xy_strings[0])
			y, _ := strconv.Atoi(xy_strings[1])
			if label_numbers[0] == "Button A" {
				button_a = XY{x, y}
			} else {
				button_b = XY{x, y}
			}
		} else if strings.HasPrefix(line, "Prize: ") {
			label_numbers := strings.Split(line, ": ")
			xy_string := strings.ReplaceAll(label_numbers[1], "X=", "")
			xy_string = strings.ReplaceAll(xy_string, "Y=", "")
			xy_strings := strings.Split(xy_string, ", ")
			x, _ := strconv.Atoi(xy_strings[0])
			y, _ := strconv.Atoi(xy_strings[1])
			prize = XY{x + 10000000000000, y + 10000000000000}
		}
	}
	fmt.Println(result)
}
