package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// check args
	if len(os.Args) != 3 {
		fmt.Println("try: file-condenser <original file> <out file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// read file
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create out file
	newFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	// create scanner
	scanner := bufio.NewScanner(file)

	// read file line by line
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// write to file with no new lines (condensed "\n")
	_, err = newFile.WriteString(strings.Join(lines, "\\n"))
	if err != nil {
		panic(err)
	}
}
