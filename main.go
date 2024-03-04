package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Changing approach - Monday:March/04/2024")
	funcWordCounter()
}

func funcWordCounter() {
	banner := `__  _____     ______        __  ____                          
\ \/ / _ |   /_  __/____ __/ /_/ __/______ ____  ___  ___ ____
 \  / __ |    / / / -_) \ / __/\ \/ __/ _ '/ _ \/ _ \/ -_/ __/
 /_/_/ |_|   /_/  \__/_\_\\__/___/\__/\_,_/_//_/_//_/\__/_/   	`

	fmt.Println(banner)
	fmt.Println("W e l c o m e    T o    t h e    Y A    T e x t S c a n n e r")

	// bufio.Method
	// Why use bufio? (fmt.scan)s are only giving
	// 1. Buffered vs unbuffered IO https://stackoverflow.com/a/1450563/21819272
	// 2. Table Difference https://gosamples.dev/read-user-input/#:~:text=kl%20mn%20op-,Use%20the%20fmt.,the%20end%20of%20each%20line.
	// scanner := bufio.NewScanner(os.Stdin)
	// input = scanner.Text()
	// scanner.Scan()
	// input = scanner.Text()
	// fmt.Printf("The enter input is: %s\n", input)

	var input string
	var word string

	fmt.Println("\nPlease enter the string!")
	scanner := bufio.NewScanner(os.Stdin)
	input = scanner.Text()
	scanner.Scan()
	input = scanner.Text()
	fmt.Printf("The enter input is: %s\n", input)
	fmt.Printf("The input is: %s", input)
	splitInput := strings.Split(input, " ")

	// fmt method
	wordCount := 0
	for _, singleInput := range splitInput {
		if word == singleInput {
			wordCount++
		}
	}
	fmt.Printf("The KeyWord %s is found %d times\n", word, wordCount)
}
