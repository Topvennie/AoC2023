package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


func check(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func inSplice(element string, list []string) bool {
	for i := 0; i < len(list); i++ {
		if element == list[i] {
			return true
		}
	}

	return false
}

func main()  {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	sum := 0
	copies := make(map[int]int)
	lineIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstPart := strings.Split(line, ":")
		secondPart := strings.Split(firstPart[1], "|")

		winningNumbers := strings.Fields(secondPart[0])
		yourNumbers := strings.Fields(secondPart[1])

		amountOfWinningCards := 0
		for i := 0; i < len(yourNumbers); i++ {
			if inSplice(yourNumbers[i], winningNumbers) {
				amountOfWinningCards++
			}
		}

		amount, found := copies[lineIndex]
		if ! found {
			amount = 0
		}

		for i := 0; i < amountOfWinningCards; i++ {
			copies[lineIndex + 1 + i] += amount + 1
		}

		sum += amount + 1
		lineIndex++
	}

	fmt.Println("Result: ", sum)
}
