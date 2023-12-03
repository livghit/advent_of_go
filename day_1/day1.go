package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Rewrite understood the problem wrong upsiiii
func Day1() {
	file, err := os.Open("./day_1/example.txt")
	if err != nil {
		panic("Could not find file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbersToAdd := []int{}

	for scanner.Scan() {
		// go trough each line of the file
		var n string
		for _, char := range scanner.Text() {
			if unicode.IsDigit(char) {
				n += string(char)
			}
			n = strings.ReplaceAll(n, " ", "")
		}
		number, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		numbersToAdd = append(numbersToAdd, number)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sum := 0
	for _, number := range numbersToAdd {
		fmt.Println("adding ", number)
		sum += number
		fmt.Println("Actual sum ", sum)
	}

	fmt.Println("Final sum ", sum)
}
