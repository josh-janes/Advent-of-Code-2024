package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLocations(filename string) ([]int, []int, error) {
	var firstNumbers, secondNumbers []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Split line by spaces

		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing first number: %w", err)
		}
		firstNumbers = append(firstNumbers, num1)

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing second number: %w", err)
		}
		secondNumbers = append(secondNumbers, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return firstNumbers, secondNumbers, nil
}

func sortInts(data []int) []int {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	return data
}

func calculateDistance(list1 []int, list2 []int) int {
	var totalDistance float64 = 0
	for i, _ := range list1 {
		totalDistance += (math.Abs(float64(list1[i] - list2[i])))
	}
	return int(totalDistance)
}

func calculateFreqCount(list2 []int) map[int]int {
	freqCount := make(map[int]int)

	for _, element := range list2 {
		freqCount[element] += 1
	}
	return freqCount
}

func calculateSimilarityScore(list1 []int, freqCount map[int]int) int {
	var similarityScore int = 0
	for _, element := range list1 {
		similarityScore += element * freqCount[element]
	}

	return similarityScore
}
func main() {
	fn, sn, err := readLocations("locations.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fn = sortInts(fn)
	sn = sortInts(sn)
	fmt.Println("First Numbers:", fn)
	fmt.Println("Second Numbers:", sn)

	fmt.Println("Total distance: ", calculateDistance(fn, sn))

	freqCount := calculateFreqCount(sn)
	similarity := calculateSimilarityScore(fn, freqCount)
	fmt.Println("Similarity score: ", similarity)
}
