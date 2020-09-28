package main

import (
	"fmt"

	"github.com/finfini/go-sdk/ekyc"
)

func main() {
	var threshold = 80.0 // threshold to reach before comparison marked as "Passed"
	var result ekyc.Similarity

	var showCase = []struct {
		first  string
		second string
	}{
		{
			first:  "widnyana putra",
			second: "widnyani putri",
		},
		{
			first:  "abcde",
			second: "bcdef",
		},
		{
			first:  "widnyana",
			second: "putra",
		},
	}

	for _, t := range showCase {
		result = ekyc.StringSimilarity(t.first, t.second, threshold)

		fmt.Printf("Comparing \"%s\" with \"%s\"; Ratio to reach: %.02f\n", t.first, t.second, threshold)
		fmt.Printf("\tRatio: %f\n", result.Ratio)
		fmt.Printf("\tPassed: %v\n", result.Passed)
		fmt.Println("====================================\n")
	}

}
