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
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		handleError(err)
	}
	defer file.Close()

	processFile(file, *byteCount, *lineCount, *wordCount, *charCount)
}

func processFile(file *os.File, byteCount, lineCount, wordCount, charCount bool) {
	switch {
	case byteCount:
		fmt.Printf("%v %v\n", byteCounter(file), file.Name())
	case lineCount:
		fmt.Printf("%v %v\n", lineCounter(file), file.Name())
	case wordCount:
		fmt.Printf("%v %v\n", wordCounter(file), file.Name())
	case charCount:
		fmt.Printf("%v %v\n", charCounter(file), file.Name())
	default:
		fmt.Printf("%v %v %v %v\n", lineCounter(file), wordCounter(file), byteCounter(file), file.Name())
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
