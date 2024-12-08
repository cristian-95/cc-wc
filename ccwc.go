package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	byteCount := flag.Bool("c", false, "Display the byte count of a .txt file.")
	args := os.Args
	if len(args) <= 2 {
		info()
		os.Exit(0)
	} else {
		flag.Parse()
		file, err := os.ReadFile(args[2])
		if err != nil {
			log.Fatal("error reading file")
		}
		if *byteCount {
			counter := 0
			counter += len(string(file))
			fmt.Printf("%v %v\n", counter, args[2])

		} else {
			fmt.Printf("%v\n", args[2])
		}
	}
}

func info() {
	fmt.Println("Usage:\n$ go build ccwc.go\n$ ./ccwc [-flags] [your-file].txt")
}
