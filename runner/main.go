package main

import (
	"fmt"
	"os"
)

func main() {
	InitFlags()
	inputFilePath := "input.txt"
	if len(os.Args) > 1 {
		inputFilePath = os.Args[1]
	}

	graph, err := parseInput(inputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %s\n", err)
		os.Exit(1)
	}

	order := LearningRoadmap(graph)

	for _, skill := range order {
		fmt.Printf("%s\n", skill.Name)
	}

}
