package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	byteCount := flag.Bool("c", false, "Display the byte count of a .txt file.")
	lineCount := flag.Bool("l", false, "Display the line count of a .txt file.")
	wordCount := flag.Bool("w", false, "Display the word count of a .txt file.")
	charCount := flag.Bool("m", false, "Display the char count of a .txt file.")
	args := os.Args

	if len(args) <= 2 {
		info()
		os.Exit(1)
	} else {
		flag.Parse()
		file, err := os.Open(args[2])
		if err != nil {
			log.Fatal("error reading file")
			os.Exit(1)
		}
		defer file.Close()

		if *byteCount {
			fmt.Printf("%v %v\n", byteCounter(file), args[2])
		}
		if *lineCount {
			fmt.Printf("%v %v\n", lineCounter(file), args[2])
		}
		if *wordCount {
			fmt.Printf("%v %v\n", wordCounter(file), args[2])
		}
		if *charCount {
			fmt.Printf("%v %v\n", charCounter(file), args[2])
		}
	}
}

func byteCounter(file *os.File) int {
	data, err := io.ReadAll(file)

	if err != nil {
		panic(err)
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
		panic(err)
	}
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
		panic(err)
	}
	return wordCount
}

func lineCounter(file *os.File) int {
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lineCount
}

func info() {
	fmt.Println("Usage:\n$ go build ccwc.go\n$ ./ccwc [-flags] [your-file].txt")
}
