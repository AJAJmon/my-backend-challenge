package services

import (
	context "context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
)

type pieFireDireServer struct {
}

func NewPieFireDireServer() PieFireDireServer {
	return pieFireDireServer{}
}

func (pieFireDireServer) mustEmbedUnimplementedPieFireDireServer() {

}

func (pieFireDireServer) Summary(ctx context.Context, req *BeefRequest) (*BeefResponse, error) {
	// Call the API to fetch data from the specified URL
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		log.Fatalf("Failed to call API: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Convert byte slice to string
	text := string(body)

	// Count the specified words in the text
	count := countWords(text)

	// Create the BeefResponse struct with counted values
	res := &BeefResponse{
		Beef: &Beef{
			TBone:    int32(count["t-bone"]),
			Fatback:  int32(count["fatback"]),
			Pastrami: int32(count["pastrami"]),
			Pork:     int32(count["pork"]),
			Meatloaf: int32(count["meatloaf"]),
			Jowl:     int32(count["jowl"]),
			Enim:     int32(count["enim"]),
			Bresaola: int32(count["bresaola"]),
		},
	}

	// Convert the response struct to JSON
	jsonData, err := protojson.Marshal(res) // Use protojson.Marshal for protobuf structs
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Print the JSON response
	fmt.Printf("Response JSON: %s\n", jsonData)

	return res, nil
}

// Function to count specified words in the text
func countWords(text string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(strings.ToLower(text)) // Split text into words and convert to lowercase

	// Specify the words to count
	targetWords := []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}

	// Count the specified words
	for _, word := range words {
		if contains(targetWords, word) {
			count[word]++
		}
	}

	return count
}

// Function to check if a word is in the slice
func contains(words []string, target string) bool {
	for _, word := range words {
		if word == target {
			return true
		}
	}
	return false
}
