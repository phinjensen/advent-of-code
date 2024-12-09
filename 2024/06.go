package main

import (
	"bufio"
	"fmt"
	"os"
)

func InitMap() (int, int, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	map_ := make([]string, 0)
	i := 0
	r, c := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		map_ = append(map_, line)
		for j, char := range line {
			if char == rune('^') {
				r = i
				c = j
			}
		}
		i++
	}
	return r, c, map_
}

func NextPosition(direction string, r int, c int) (int, int) {
	next_r, next_c := r, c
	if direction == "n" {
		next_r--
	} else if direction == "e" {
		next_c++
	} else if direction == "s" {
		next_r++
	} else if direction == "w" {
		next_c--
	}
	return next_r, next_c
}

func NextDirection(direction string) string {
	if direction == "n" {
		return "e"
	} else if direction == "e" {
		return "s"
	} else if direction == "s" {
		return "w"
	} else if direction == "w" {
		return "n"
	}
	return "INVALID"
}

func OffMap(r int, c int, map_ []string) bool {
	return r < 0 || r >= len(map_) || c < 0 || c >= len(map_[0])
}

func Day6_1() {
	r, c, map_ := InitMap()
	visited := make(map[XY]bool)
	result := 0
	direction := "n"
	for {
		if !visited[XY{r, c}] {
			result += 1
			visited[XY{r, c}] = true
		}
		next_r, next_c := NextPosition(direction, r, c)
		if OffMap(next_r, next_c, map_) {
			break
		}
		if map_[next_r][next_c] == byte('#') {
			next_r, next_c = r, c
			direction = NextDirection(direction)
		}
		r, c = next_r, next_c
	}
	fmt.Println(result)
}

type XYD struct {
	X         int
	Y         int
	Direction string
}

func Day6_2() {
	r, c, map_ := InitMap()
	result := 0
	direction := "n"
	successful_obstacles := make(map[XY]bool)
	visited := make(map[XY]bool)
	for {
		visited[XY{r, c}] = true
		next_r, next_c := NextPosition(direction, r, c)
		if next_r < 0 || next_r >= len(map_) || next_c < 0 || next_c >= len(map_[0]) {
			break
		}
		if map_[next_r][next_c] == byte('#') {
			next_r, next_c = r, c
			direction = NextDirection(direction)
		} else if !successful_obstacles[XY{next_r, next_c}] && !visited[XY{next_r, next_c}] {
			// Send a "ray" to the right
			original := map_[next_r]
			map_[next_r] = map_[next_r][:next_c] + "#" + map_[next_r][next_c+1:]
			ray_r, ray_c := r, c
			ray_direction := NextDirection(direction)
			ray_visited := make(map[XYD]bool)
			for {
				prev_ray_r, prev_ray_c := ray_r, ray_c
				ray_r, ray_c = NextPosition(ray_direction, ray_r, ray_c)
				if OffMap(ray_r, ray_c, map_) {
					break
				}
				if ray_visited[XYD{ray_r, ray_c, ray_direction}] {
					fmt.Println(next_r, next_c)
					successful_obstacles[XY{next_r, next_c}] = true
					result += 1
					break
				} else if map_[ray_r][ray_c] == byte('#') {
					ray_visited[XYD{ray_r, ray_c, ray_direction}] = true
					ray_r, ray_c = prev_ray_r, prev_ray_c
					ray_direction = NextDirection(ray_direction)
				}
			}
			map_[next_r] = original
		}
		r, c = next_r, next_c
	}
	fmt.Println(len(successful_obstacles))
	fmt.Println(result)
}
