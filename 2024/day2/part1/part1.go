package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseIntegers(line string) []int {
	numStrings := strings.Fields(line)

	var answer []int
	for i := range len(numStrings) {
		num, err := strconv.Atoi(numStrings[i])
		if err != nil {
			log.Fatal(err)
		} else {
			answer = append(answer, num)
		}
	}

	return answer
}

func isSafe(arr []int) bool {
	var answer bool = true
	initialDifference := arr[0] - arr[1]
	var count int = 0
	if initialDifference >= 0 {
		for i := range len(arr) - 1 {
			difference := arr[i] - arr[i+1]
			if difference <= 3 && difference > 0 {
				count++
			}
		}
	} else {
		for i := range len(arr) - 1 {
			difference := arr[i] - arr[i+1]
			if difference < 0 && difference >= -3 {
				count++
			}
		}
	}

	if count != len(arr)-1 {
		answer = false
	}

	return answer
}

func main() {
	// Get input from the file
	input, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// Go through the input one by one
	scanner := bufio.NewScanner(input)

	var count int = 0

	for scanner.Scan() {
		line := scanner.Text()
		arr := parseIntegers(line)
		if isSafe(arr) {
			count++
		}
	}

    fmt.Println(count)

}
