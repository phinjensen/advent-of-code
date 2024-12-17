package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Day11_1() {
	input_bytes, _ := io.ReadAll(os.Stdin)
	input := strings.TrimSpace(string(input_bytes))
	stones := strings.Split(input, " ")
	for i := range 75 {
		fmt.Println(i)
		new_stones := make([]string, 0)
		for _, stone := range stones {
			if len(stone)%2 == 0 {
				a, b := stone[:len(stone)/2], stone[len(stone)/2:]
				a_i, _ := strconv.Atoi(a)
				b_i, _ := strconv.Atoi(b)
				new_stones = append(new_stones, strconv.Itoa(a_i))
				new_stones = append(new_stones, strconv.Itoa(b_i))
			} else if stone == "0" {
				new_stones = append(new_stones, "1")
			} else {
				stone_value, _ := strconv.Atoi(stone)
				stone_value *= 2024
				new_stones = append(new_stones, strconv.Itoa(stone_value))
			}
		}
		stones = new_stones
	}
	fmt.Println(len(stones))
}

type StoneKey struct {
	stone      string
	steps_left int
}

func ResultForStone(stone string, steps_left int, memo map[StoneKey]int) int {
	result := 0
	key := StoneKey{stone, steps_left}
	if steps_left > 0 {
		memo_value, found := memo[key]
		if found {
			return memo_value
		} else if len(stone)%2 == 0 {
			a, b := stone[:len(stone)/2], stone[len(stone)/2:]
			a_i, _ := strconv.Atoi(a)
			b_i, _ := strconv.Atoi(b)
			result += ResultForStone(strconv.Itoa(a_i), steps_left-1, memo)
			result += ResultForStone(strconv.Itoa(b_i), steps_left-1, memo)
		} else if stone == "0" {
			result += ResultForStone("1", steps_left-1, memo)
		} else {
			stone_value, _ := strconv.Atoi(stone)
			stone_value *= 2024
			result += ResultForStone(strconv.Itoa(stone_value), steps_left-1, memo)
		}
	} else {
		result = 1
	}
	memo[key] = result
	return result
}

func Day11_2() {
	input_bytes, _ := io.ReadAll(os.Stdin)
	input := strings.TrimSpace(string(input_bytes))
	stones := strings.Split(input, " ")
	memo := make(map[StoneKey]int)
	result := 0
	for _, stone := range stones {
		this_result := ResultForStone(stone, 75, memo)
		result += this_result
	}
	fmt.Println(result)
}
