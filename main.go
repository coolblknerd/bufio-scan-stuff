package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanFile := bufio.NewScanner(file)
	scanFile.Split(bufio.ScanWords)

	var words []string

	for scanFile.Scan() {
		words = append(words, scanFile.Text())
	}

	// This prints out all the words
	fmt.Println("word list:")
	for _, word := range words {
		fmt.Println(word)
	}
}
