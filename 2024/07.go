package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsTestValueValid(target int, operands []int) bool {
	if len(operands) == 2 {
		if operands[0]*operands[1] == target || operands[0]+operands[1] == target {
			return true
		}
	} else {
		multiplied := make([]int, len(operands))
		copy(multiplied, operands)
		multiplied[1] = multiplied[0] * multiplied[1]
		if IsTestValueValid(target, multiplied[1:]) {
			return true
		}
		added := make([]int, len(operands))
		copy(added, operands)
		added[1] = added[0] + added[1]
		if IsTestValueValid(target, added[1:]) {
			return true
		}
	}
	return false
}

func Day7_1() {
	scanner := bufio.NewScanner(os.Stdin)
	result := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		test_value, _ := strconv.Atoi(split[0])
		operands := make([]int, 0)
		for _, v := range strings.Split(split[1], " ") {
			operand, _ := strconv.Atoi(v)
			operands = append(operands, operand)
		}
		if IsTestValueValid(test_value, operands) {
			result += test_value
		}
	}
	fmt.Println(result)
}

func ConcatIntegers(a int, b int) int {
	s := strconv.Itoa(a) + strconv.Itoa(b)
	result, _ := strconv.Atoi(s)
	return result
}

func IsTestValueValidWithConcat(target int, operands []int) bool {
	if len(operands) == 2 {
		concat := ConcatIntegers(operands[0], operands[1])
		if operands[0]*operands[1] == target || operands[0]+operands[1] == target || concat == target {
			return true
		}
	} else {
		multiplied := make([]int, len(operands))
		copy(multiplied, operands)
		multiplied[1] = multiplied[0] * multiplied[1]
		if IsTestValueValidWithConcat(target, multiplied[1:]) {
			return true
		}
		added := make([]int, len(operands))
		copy(added, operands)
		added[1] = added[0] + added[1]
		if IsTestValueValidWithConcat(target, added[1:]) {
			return true
		}
		concated := make([]int, len(operands))
		copy(concated, operands)
		concated[1] = ConcatIntegers(concated[0], concated[1])
		if IsTestValueValidWithConcat(target, concated[1:]) {
			return true
		}
	}
	return false
}

func Day7_2() {
	scanner := bufio.NewScanner(os.Stdin)
	result := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		test_value, _ := strconv.Atoi(split[0])
		operands := make([]int, 0)
		for _, v := range strings.Split(split[1], " ") {
			operand, _ := strconv.Atoi(v)
			operands = append(operands, operand)
		}
		if IsTestValueValidWithConcat(test_value, operands) {
			result += test_value
		}
	}
	fmt.Println(result)
}
