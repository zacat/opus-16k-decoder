package main

import (
	"fmt"
	"io"
	"os"
	"github.com/yourusername/yourrepo/decoder" // Replace with actual import path
)

func main() {
	inputFile, err := os.Open("input.opus")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("output.wav")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	decoder := decoder.NewDecoder() // Adjust according to your decoder's API

	if err := decoder.StreamDecode(inputFile, outputFile); err != nil {
		fmt.Println("Error during streaming decode:", err)
	}
}