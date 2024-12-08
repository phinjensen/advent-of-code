package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func GetRules(scanner *bufio.Scanner) map[int][]int {
	rules := make(map[int][]int)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules[before] = append(rules[before], after)
	}

	return rules
}

func GetPages(scanner *bufio.Scanner) [][]int {
	result := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, ",")
		pages := make([]int, len(line_split))
		for i, n := range line_split {
			pages[i], _ = strconv.Atoi(n)
		}
		result = append(result, pages)
	}
	return result
}

func IsValid(pages []int, rules map[int][]int) bool {
	valid := true
	for i, v := range pages {
		followers := rules[v]
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(followers, pages[j]) {
				valid = false
				break
			}
		}
		if !valid {
			break
		}
	}
	return valid
}

func Day5_1() {
	scanner := bufio.NewScanner(os.Stdin)
	rules := GetRules(scanner)
	pages_list := GetPages(scanner)

	result := 0
	for _, pages := range pages_list {
		if IsValid(pages, rules) {
			result += pages[len(pages)/2]
		}
	}

	fmt.Println(result)
}

// Doesn't preserve order
func Remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func InvertRules(rules map[int][]int) map[int][]int {
	result := make(map[int][]int)
	for k, v := range rules {
		for _, vertex := range v {
			result[vertex] = append(result[vertex], k)
		}
	}
	return result
}

func Day5_2() {
	scanner := bufio.NewScanner(os.Stdin)
	rules_out := GetRules(scanner)
	rules := InvertRules(rules_out)
	pages_list := GetPages(scanner)

	result := 0
	for _, pages := range pages_list {
		if IsValid(pages, rules_out) {
			continue
		}
		local_rules := make(map[int][]int)
		for k, v := range rules {
			if slices.Contains(pages, k) {
				local_v := make([]int, 0)
				for _, p := range v {
					if slices.Contains(pages, p) {
						local_v = append(local_v, p)
					}
				}
				local_rules[k] = local_v
			}
		}

		ordered := make([]int, 0)

		for len(pages) > 0 {
			for i := range len(pages) {
				v := pages[i]
				if len(local_rules[v]) == 0 {
					delete(local_rules, v)
					for j, l := range local_rules {
						v_i := slices.Index(l, v)
						if v_i >= -1 {
							local_rules[j] = Remove(l, v_i)
						}
					}
					ordered = append(ordered, v)
					pages[i] = pages[len(pages)-1]
					pages = pages[:len(pages)-1]
					break
				}
			}
		}

		result += ordered[len(ordered)/2]
	}
	fmt.Println(result)
}
