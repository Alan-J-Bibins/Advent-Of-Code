package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quicksort(arr []int, low int, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		quicksort(arr, low, partitionIndex-1)
		quicksort(arr, partitionIndex+1, high)
	}
}

func parseIntegers(line string) (int, int, error) {
	parts := strings.Fields(line)
	var num1 int = 0
	var num2 int = 0
	var error error = nil
	if n, err := strconv.Atoi(parts[0]); err == nil {
		num1 = n
	} else {
		error = err
	}

	if n, err := strconv.Atoi(parts[1]); err == nil {
		num2 = n
	} else {
		error = err
	}

	return num1, num2, error
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	} else {
		return num
	}
}

func main() {
	// Getting the input from a file input.txt
	input, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// taking input numbers and putting them in two arrays
	var arr1 []int
	var arr2 []int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		num1, num2, err := parseIntegers(line)
		if err != nil {
			log.Fatal(err)
		}

		arr1 = append(arr1, num1)
		arr2 = append(arr2, num2)
	}

	// sorting the two arrays
	quicksort(arr1, 0, len(arr1)-1)
	quicksort(arr2, 0, len(arr2)-1)

	var difference []int
	// getting difference of sorted arrays index-wise
	for i := range len(arr1) {
		difference = append(difference, abs(arr2[i]-arr1[i]))
	}

	// getting the sum of the difference slice
	sum := 0
	for i := range len(difference) {
		sum += difference[i]
	}

	for i := range len(difference) {
		fmt.Printf("%d - %d\n", arr1[i], arr2[i])
	}

	fmt.Println("Answer = ", sum)
}
