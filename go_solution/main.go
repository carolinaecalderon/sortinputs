package main

import (
	"fmt"
	"os"
)

func main() {
	// check for correct input
	if len(os.Args) != 3 {
		fmt.Println("to run this program, `go run main.go <input_directory> <output_file>`")
		os.Exit(1)
	}

	inputDir := os.Args[1]
	outputFile := os.Args[2]

	// consolidate & sort the data into one output file
	if err := consolidateData(inputDir, outputFile); err != nil {
		fmt.Printf("error consolidating data: %s ", err)
		os.Exit(1)
	}
}
