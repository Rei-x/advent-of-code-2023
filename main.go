package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	text := strings.Split(string(file), "\n")

	sum := 0

	for _, line := range text {
		numbers := []int{}

		for _, character := range line {
			if character >= '0' && character <= '9' {
				numbers = append(numbers, int(character-'0'))
			}
		}

		if len(numbers) > 0 {

			firstDigit := numbers[0]
			lastDigit := numbers[len(numbers)-1]
			sum += firstDigit*10 + lastDigit
		}

	}

	fmt.Println("part one", sum)
}

func partTwo() {
	stringToNumber := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	text := strings.Split(string(file), "\n")

	sum := 0

	for _, line := range text {
		firstIndex := len(line)
		lastIndex := -1

		firstDigit := 0
		lastDigit := 0

		for key, value := range stringToNumber {
			firstFoundIndex := strings.Index(line, key)
			lastFoundIndex := strings.LastIndex(line, key)

			if firstFoundIndex == -1 && lastFoundIndex == -1 {
				continue
			}

			if firstFoundIndex < firstIndex {
				firstIndex = firstFoundIndex
				firstDigit = value
			}

			if lastFoundIndex > lastIndex {
				lastIndex = lastFoundIndex
				lastDigit = value
			}
		}

		sum += firstDigit*10 + lastDigit
	}

	fmt.Println("part two", sum)
}
