package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open the JSON file
	file, err := os.Open("hard.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the JSON data into a [][]int slice
	var data [][]int
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the result
	fmt.Println(findMaxPathSum(data))
}

// Function to find the maximum path sum
func findMaxPathSum(triangle [][]int) int {
	// Iterate from the second last row to the top
	triangleChoosed := 0
	triangleBottomLeft := 0
	triangleBottomRight := 0
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// Choose the maximum sum path from the two possible paths below
			triangleBottomLeft = triangle[i+1][j]
			triangleBottomRight = triangle[i+1][j+1]

			if triangleBottomLeft > triangleBottomRight {
				triangleChoosed = triangleBottomLeft
			} else {
				triangleChoosed = triangleBottomRight
			}
			triangle[i][j] += triangleChoosed
		}
	}
	// The maximum sum is at the top of the triangle
	return triangle[0][0]
}
