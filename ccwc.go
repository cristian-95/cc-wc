package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	byteCount := flag.Bool("c", false, "Display the byte count of a .txt file.")
	lineCount := flag.Bool("l", false, "Display the line count of a .txt file.")
	wordCount := flag.Bool("w", false, "Display the word count of a .txt file.")
	charCount := flag.Bool("m", false, "Display the char count of a .txt file.")
	flag.Parse()

	var file *os.File
	var err error
	name := ""

	if len(flag.Args()) > 0 {
		file, err = os.Open(flag.Arg(0))
		if err != nil {
			handleError(err)
		}
		name = file.Name()
	} else {
		file = os.Stdin
	}

	defer file.Close()

	processFile(file, name, *byteCount, *lineCount, *wordCount, *charCount)
}

func processFile(file *os.File, fileName string, byteCount, lineCount, wordCount, charCount bool) {
	switch {
	case byteCount:
		fmt.Printf("%v %v\n", byteCounter(file), fileName)
	case lineCount:
		fmt.Printf("%v %v\n", lineCounter(file), fileName)
	case wordCount:
		fmt.Printf("%v %v\n", wordCounter(file), fileName)
	case charCount:
		fmt.Printf("%v %v\n", charCounter(file), fileName)
	default:
		fmt.Printf("%v %v %v %v\n", lineCounter(file), wordCounter(file), byteCounter(file), fileName)
	}
}

func byteCounter(file *os.File) int {
	data, err := io.ReadAll(file)

	if err != nil {
		handleError(err)
	}
	return len(data)
}

func wordCounter(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		handleError(err)
	}

	file.Seek(0, io.SeekStart)
	return wordCount
}

func charCounter(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		handleError(err)
	}

	file.Seek(0, io.SeekStart)
	return wordCount
}

func lineCounter(file *os.File) int {
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		handleError(err)
	}

	file.Seek(0, io.SeekStart)
	return lineCount
}

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
