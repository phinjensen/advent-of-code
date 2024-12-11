package main

import (
	"fmt"
	"io"
	"os"
  "slices"
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
      checksum += (i-1) * id_number
      //fmt.Println("i:", i, "number:", id_number, "current_space:", current_space, "checksum:", checksum)
    }
	}
	fmt.Println(checksum)
}

func Day9_2() {
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
	in_space := false
  current_space := 0
	checksum := 0
  for {
    if in_space {
      for j, s := range slices.Backward(files) {
        if s <= space[current_space] {
          space[current_space] -= s
          for k := 0; k < s; k++ {
            checksum += i * j
            i++
          }
        }
      }
      if space[current_space] == 0 {
        current_space += 1
        in_space = false
      }
    } else {
			if files[front] == 0 {
        //fmt.Println("end of number", front)
				front += 1
				in_space = true
			} else {
				files[front] -= 1
        checksum += i * front
				i++
			}
		}
  }
}
