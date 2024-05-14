package main

import (
	"fmt"
)

func processString(input string, num []int, index int) int {
	char := input[index]
	switch char {
	case 'L':
		if num[index] <= num[index+1] {
			num[index]++
		}
	case 'R':
		if num[index] >= num[index+1] {
			num[index+1]++
		}
	}
	return checkAnsNum(input, num)
}

func checkAnsNum(input string, num []int) int {
	countTrue := 0
	inputLength := len(input)
	for i := 0; i < len(input); i++ {
		char := input[i]
		switch char {
		case 'L':
			if num[i] > num[i+1] {
				countTrue++
			} else {
				return processString(input, num, i)
			}
		case 'R':
			if num[i] < num[i+1] {
				countTrue++
			} else {
				return processString(input, num, i)
			}
		}
	}

	// Check how many equal
	numOfEqual := 0
	for _, equal := range input {
		if equal == '=' {
			numOfEqual++
		}
	}

	// When codition is true
	if countTrue == inputLength-numOfEqual {
		// Check case '='
		for i := 0; i < inputLength; i++ {
			char := input[i]
			switch char {
			case '=':
				if num[i] > num[i+1] {
					num[i+1] = num[i]
				} else {
					num[i] = num[i+1]
				}
			}
		}
		// Output
		for _, ansNum := range num {
			fmt.Print(ansNum)
		}
		return 1
	}
	return 0
}

func main() {
	var input string
	fmt.Scanln(&input)
	num := make([]int, len(input)+1)
	processString(input, num, 0)
}

//	Test case
/*
	LLRR=
	==RLL
	=LLRR
	RRL=R
	LLLRR
	LLLLL
	RRRRR
	=====
*/
/*
	210122
	000210
	221012
	012001
	321012
	543210
	012345
	000000
*/
