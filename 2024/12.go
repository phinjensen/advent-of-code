package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (a XY) Plus(b XY) XY {
	return XY{a.X + b.X, a.Y + b.Y}
}

func Day12_1() {
	scanner := bufio.NewScanner(os.Stdin)
	garden := make([][]string, 0)
	for scanner.Scan() {
		garden = append(
			garden,
			strings.Split(
				strings.TrimSpace(scanner.Text()),
				"",
			),
		)
	}
	visited := make([][]bool, len(garden))
	for i := range garden {
		visited[i] = make([]bool, len(garden[i]))
	}
	directions := [4]XY{
		{-1, 0}, // n
		{0, 1},  // e
		{1, 0},  // s
		{0, -1}, // w
	}
	result := 0
	for r, row := range garden {
		for c, plant := range row {
			if visited[r][c] {
				continue
			}
			perimeter := 0
			area := 0
			stack := make([]XY, 0)
			stack = append(stack, XY{r, c})
			for len(stack) > 0 {
				l := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				rn, rc := l.X, l.Y
				if rn < 0 || rn >= len(garden) || rc < 0 || rc >= len(row) {
					perimeter += 1
				} else if garden[rn][rc] != plant {
					perimeter += 1
				} else if !visited[rn][rc] {
					area += 1
					visited[rn][rc] = true
					for _, direction := range directions {
						stack = append(stack, l.Plus(direction))
					}
				}
			}
			result += area * perimeter
		}
	}
	fmt.Println(result)
}

type Movement struct {
	location  XY
	direction XY
}

func Day12_2() {
	scanner := bufio.NewScanner(os.Stdin)
	garden := make([][]string, 0)
	for scanner.Scan() {
		garden = append(
			garden,
			strings.Split(
				strings.TrimSpace(scanner.Text()),
				"",
			),
		)
	}
	visited := make([][]bool, len(garden))
	for i := range garden {
		visited[i] = make([]bool, len(garden[i]))
	}
	directions := [4]XY{
		{-1, 0}, // n
		{0, 1},  // e
		{1, 0},  // s
		{0, -1}, // w
	}
	result := 0
	for r, row := range garden {
		for c, plant := range row {
			if visited[r][c] {
				continue
			}
			perimeters := make(map[Movement]bool)
			area := 0
			stack := make([]Movement, 0)
			stack = append(stack, Movement{XY{r, c}, XY{0, 0}})
			for len(stack) > 0 {
				m := stack[len(stack)-1]
				l := m.location.Plus(m.direction)
				stack = stack[:len(stack)-1]
				rn, rc := l.X, l.Y
				if (rn < 0 || rn >= len(garden) || rc < 0 || rc >= len(row)) || garden[rn][rc] != plant {
					perimeters[m] = true
				} else if !visited[rn][rc] {
					area += 1
					visited[rn][rc] = true
					for _, direction := range directions {
						stack = append(stack, Movement{l, direction})
					}
				}
			}
			perimeter := 0
			for k, _ := range perimeters {
				direction := k.direction
				k2 := k
				var d1, d2 XY
				if direction == (XY{-1, 0}) || direction == (XY{1, 0}) {
					d1 = XY{0, -1}
					d2 = XY{0, 1}
				} else {
					d1 = XY{-1, 0}
					d2 = XY{1, 0}
				}
				perimeter += 1
				delete(perimeters, k)
				k2.location = k2.location.Plus(d1)
				for perimeters[k2] {
					delete(perimeters, k2)
					k2.location = k2.location.Plus(d1)
				}
				k2.location = k.location.Plus(d2)
				for perimeters[k2] {
					delete(perimeters, k2)
					k2.location = k2.location.Plus(d2)
				}
			}
			fmt.Println(plant, area, perimeter)
			result += area * perimeter
		}
	}
	fmt.Println(result)
}
