package main

import (
	"fmt"
)

func processString(input string) {
	inputLength := len(input)
	num := make([]int, len(input)+1)
	left, right := true, true
	for left || right {
		for i := 0; i < inputLength; i++ {
			char := input[i]
			switch char {
			case 'L':
				if num[i] > num[i+1] {
					left = false
				} else {
					num[i]++
				}
			case 'R':
				if num[i] < num[i+1] {
					right = false
				} else {
					num[i+1]++
				}
			}
		}
	}
	// Check case '='
	for i := 0; i < inputLength; i++ {
		char := input[i]
		switch char {
		case '=':
			if i != inputLength-1 {
				num[i] = num[i+1]
			} else {
				num[i+1] = num[i]
			}
		}
	}
	// Output
	for _, ansNum := range num {
		fmt.Print(ansNum)
	}
}

func main() {
	var input string
	fmt.Scanln(&input)

	// Process the input string
	processString(input)
}

//	Test case
/*
	LLRR=
	==RLL
	=LLRR
	RRL=R
*/
/*
	210122
	000210
	221012
	012001
*/
