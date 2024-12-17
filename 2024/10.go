package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func LocationIsValid(location XY, topo_map []string) bool {
	return location.X >= 0 && location.X < len(topo_map) && location.Y >= 0 && location.Y < len(topo_map[0])
}

func FindTrails(locations []XY, topo_map []string, visited map[XY]bool) int {
	result := 0
	for _, l := range locations {
		if visited != nil {
			visited[l] = true
		}
		height := topo_map[l.X][l.Y]
		if height == '9' {
			result += 1
		} else {
			next_locations := make([]XY, 0)
			for _, delta := range [4][2]int{
				{0, 1},
				{1, 0},
				{0, -1},
				{-1, 0},
			} {
				next := XY{l.X + delta[0], l.Y + delta[1]}
				if LocationIsValid(next, topo_map) && topo_map[next.X][next.Y] == height+1 && (visited == nil || !visited[next]) {
					next_locations = append(next_locations, next)
				}
			}
			result += FindTrails(next_locations, topo_map, visited)
		}
	}
	return result
}

func Day10_1() {
	input, _ := io.ReadAll(os.Stdin)
	topo_map := strings.Split(strings.TrimSpace(string(input)), "\n")
	result := 0
	for r, line := range topo_map {
		for c, char := range line {
			if char == '0' {
				visited := make(map[XY]bool)
				start := make([]XY, 0)
				start = append(start, XY{r, c})
				score := FindTrails(start, topo_map, visited)
				fmt.Println(r, c, score)
				result += score
			}
		}
	}
	fmt.Println(result)
}

func Day10_2() {
	input, _ := io.ReadAll(os.Stdin)
	topo_map := strings.Split(strings.TrimSpace(string(input)), "\n")
	result := 0
	for r, line := range topo_map {
		for c, char := range line {
			if char == '0' {
				start := make([]XY, 0)
				start = append(start, XY{r, c})
				score := FindTrails(start, topo_map, nil)
				fmt.Println(r, c, score)
				result += score
			}
		}
	}
	fmt.Println(result)
}
