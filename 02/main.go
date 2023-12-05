package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.ReadFile("input.txt")
	numberOfRedCubes := 12
	numberOfGreenCubes := 13
	numberOfBlueCubes := 14
	if err != nil {
		log.Fatal(err)
		return
	}

	text := strings.Split(string(file), "\n")

	sum := 0

	for _, line := range text {
		splittedByColon := strings.Split(line, ":")
		game := strings.Split(splittedByColon[0], " ")[1]

		gameId, err := strconv.ParseInt(game, 10, 8)

		if err != nil {
			log.Fatal(err)
			return
		}
		games := strings.Split(splittedByColon[1], ";")

		isPossible := true

		for _, singleGame := range games {
			trimmedGame := strings.Trim(singleGame, " ")

			allColors := strings.Split(trimmedGame, ",")

			for _, currentColor := range allColors {
				trimmedColor := strings.Split(strings.Trim(currentColor, " "), " ")

				color := trimmedColor[1]
				number := trimmedColor[0]

				parsedNumber, err := strconv.ParseInt(number, 10, 0)

				if err != nil {
					log.Fatal(err)
					return
				}

				switch color {
				case "red":
					if int(parsedNumber) > numberOfRedCubes {
						isPossible = false
					}
				case "blue":
					if int(parsedNumber) > numberOfBlueCubes {
						isPossible = false
					}
				case "green":
					if int(parsedNumber) > numberOfGreenCubes {
						isPossible = false
					}
				}
			}

		}

		if isPossible {
			sum += int(gameId)
		}
	}

	fmt.Println("part1", sum)
}

func partTwo() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	text := strings.Split(string(file), "\n")

	sum := 0

	for _, line := range text {
		splittedByColon := strings.Split(line, ":")

		if err != nil {
			log.Fatal(err)
			return
		}
		games := strings.Split(splittedByColon[1], ";")

		minimumNumberOfRedCubes := 0
		minimumNumberOfBlueCubes := 0
		minimumNumberOfGreenCubes := 0

		for _, singleGame := range games {
			trimmedGame := strings.Trim(singleGame, " ")

			allColors := strings.Split(trimmedGame, ",")

			for _, currentColor := range allColors {
				trimmedColor := strings.Split(strings.Trim(currentColor, " "), " ")

				color := trimmedColor[1]
				number := trimmedColor[0]

				parsedNumber, err := strconv.ParseInt(number, 10, 0)

				if err != nil {
					log.Fatal(err)
					return
				}

				switch color {
				case "red":
					minimumNumberOfRedCubes = max(minimumNumberOfRedCubes, int(parsedNumber))
				case "blue":
					minimumNumberOfBlueCubes = max(minimumNumberOfBlueCubes, int(parsedNumber))
				case "green":
					minimumNumberOfGreenCubes = max(minimumNumberOfGreenCubes, int(parsedNumber))
				}
			}
		}

		sum += minimumNumberOfBlueCubes * minimumNumberOfGreenCubes * minimumNumberOfRedCubes

	}

	fmt.Println("part2", sum)
}
