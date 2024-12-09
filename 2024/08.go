package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func BuildAntennaMap() (input []string, antenna_map map[rune][]XY, w int, h int) {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
	input = strings.Split(strings.TrimSpace(string(bytes)), "\n")
	antenna_map = make(map[rune][]XY)
	for r, line := range input {
		for c, char := range line {
			if (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
				antenna_map[char] = append(antenna_map[char], XY{r, c})
			}
		}
	}
	h = len(input)
	w = len(input[0])
	return input, antenna_map, w, h
}

func Day8_1() {
	input, antenna_map, w, h := BuildAntennaMap()
	antinodes := make(map[XY]bool, 0)
	for _, antennas := range antenna_map {
		for i, a := range antennas {
			for _, b := range antennas[i+1:] {
				// It's confusing, but X is up-down here, since we're doing row-first coordinates
				rise := a.X - b.X
				run := a.Y - b.Y
				a_x := a.X + rise
				a_y := a.Y + run
				if a_x >= 0 && a_x < h && a_y >= 0 && a_y < w {
					antinodes[XY{a_x, a_y}] = true
					input[a_x] = input[a_x][:a_y] + "#" + input[a_x][a_y+1:]
				}
				a_x = b.X - rise
				a_y = b.Y - run
				if a_x >= 0 && a_x < h && a_y >= 0 && a_y < w {
					antinodes[XY{a_x, a_y}] = true
					input[a_x] = input[a_x][:a_y] + "#" + input[a_x][a_y+1:]
				}
			}
		}
	}
	for _, line := range input {
		fmt.Println(line)
	}
	fmt.Println(len(antinodes))
}

func Day8_2() {
	input, antenna_map, w, h := BuildAntennaMap()
	antinodes := make(map[XY]bool, 0)
	for _, antennas := range antenna_map {
		for i, a := range antennas {
			for _, b := range antennas[i+1:] {
				// It's confusing, but X is up-down here, since we're doing row-first coordinates
				rise := a.X - b.X
				run := a.Y - b.Y
				a_x := a.X
				a_y := a.Y
				for a_x >= 0 && a_x < h && a_y >= 0 && a_y < w {
					antinodes[XY{a_x, a_y}] = true
					input[a_x] = input[a_x][:a_y] + "#" + input[a_x][a_y+1:]
					a_x = a_x + rise
					a_y = a_y + run
				}
				a_x = b.X
				a_y = b.Y
				for a_x >= 0 && a_x < h && a_y >= 0 && a_y < w {
					antinodes[XY{a_x, a_y}] = true
					input[a_x] = input[a_x][:a_y] + "#" + input[a_x][a_y+1:]
					a_x = a_x - rise
					a_y = a_y - run
				}
			}
		}
	}
	for _, line := range input {
		fmt.Println(line)
	}
	fmt.Println(len(antinodes))
}
