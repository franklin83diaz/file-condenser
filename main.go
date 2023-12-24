package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	// regex to replace new lines
	re := regexp.MustCompile(`\r?\n`)

	// read file line by line
	var lines []string
	for scanner.Scan() {
		// remplace regex with \n
		line := re.ReplaceAllString(scanner.Text(), "\\n")
		lines = append(lines, line)
	}

	// write to file with no new lines (condensed "\n")
	_, err = newFile.WriteString(strings.Join(lines, ""))
	if err != nil {
		panic(err)
	}
}
