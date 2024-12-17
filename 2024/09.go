package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func Day9_1() {
	input, _ := io.ReadAll(os.Stdin)
	disk_map := string(input)
	files := make([]int, 0)
	space := make([]int, 0)
	for i, c := range disk_map {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			files = append(files, n)
		} else {
			space = append(space, n)
		}
	}
	i := 0
	front := 0
	back := len(files) - 1
	in_space := false
	current_space := 0
	checksum := 0
	for front <= back {
		id_number := -1
		if in_space {
			if space[current_space] == 0 {
				//fmt.Println("end of space")
				in_space = false
				current_space += 1
			} else if files[back] == 0 {
				//fmt.Println("end of file", back)
				back -= 1
			} else {
				id_number = back
				space[current_space] -= 1
				files[back] -= 1
				i++
			}
		} else {
			if files[front] == 0 {
				//fmt.Println("end of number", front)
				front += 1
				in_space = true
			} else {
				id_number = front
				files[front] -= 1
				i++
			}
		}
		if id_number >= 0 {
			checksum += (i - 1) * id_number
			//fmt.Println("i:", i, "number:", id_number, "current_space:", current_space, "checksum:", checksum)
		}
	}
	fmt.Println(checksum)
}

type Chunk struct {
	size  int
	space bool
	moved bool
	id    int
}

func Day9_2() {
	input, _ := io.ReadAll(os.Stdin)
	disk_map_string := string(input)
	disk_map := make([]Chunk, 0)
	for i, c := range disk_map_string {
		n, _ := strconv.Atoi(string(c))
		var chunk Chunk
		if i%2 == 0 {
			chunk = Chunk{n, false, false, i / 2}
		} else {
			chunk = Chunk{n, true, false, 0}
		}
		disk_map = append(disk_map, chunk)
	}

	disk_map = disk_map[:len(disk_map)-1]

	for i := len(disk_map) - 1; i > 0; i-- {
		chunk := disk_map[i]
		if !chunk.space && !chunk.moved {
			for j := 0; j < i; j++ {
				left_chunk := disk_map[j]
				if left_chunk.space && left_chunk.size >= chunk.size {
					left_chunk.size -= chunk.size
					chunk.moved = true
					temp := make([]Chunk, 0)
					// Add up to space
					temp = append(temp, disk_map[:j]...)
					// Add chunk to beginning of space
					temp = append(temp, chunk)
					// Add space if it still has size
					if left_chunk.size > 0 {
						temp = append(temp, left_chunk)
					}
					// Add from after space to just before chunk's old position
					temp = append(temp, disk_map[j+1:i]...)
					before := &temp[len(temp)-1]
					if before.space {
						before.size += chunk.size
						if i+1 < len(disk_map) {
							after := disk_map[i+1]
							if after.space {
								before.size += after.size
							}
						}
						if i+2 < len(disk_map) {
							temp = append(temp, disk_map[i+2:]...)
						}
					} else {
						if i+1 < len(disk_map) {
							after := disk_map[i+1]
							if after.space {
								after.size += chunk.size
							} else {
								temp = append(temp, Chunk{chunk.size, true, false, 0})
							}
						}
						temp = append(temp, disk_map[i+1:]...)
					}
					disk_map = temp
					break
				}
			}
		}
	}

	i := 0
	checksum := 0
	for _, chunk := range disk_map {
		for range chunk.size {
			if chunk.space {
				fmt.Print(".")
			} else {
				fmt.Print(chunk.id)
				checksum += i * chunk.id
			}
			i++
		}
	}
	fmt.Println()
	fmt.Println(checksum)
}
