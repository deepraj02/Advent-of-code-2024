package main

import (
	"bufio"
	distance "deepraj02/aoc2024/part1"
	similarity "deepraj02/aoc2024/part2"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "inputs.txt"
	leftList, rightList := inputParser(filepath)
	totalDistance := distance.CalculateDistance(leftList, rightList)
	fmt.Printf("Total distance between lists: %d\n", totalDistance)
	similarityScore := similarity.CalculateSimilarityScore(leftList, rightList)
	fmt.Printf("Similarity score between lists: %d\n", similarityScore)

}

func inputParser(filename string) ([]int, []int) {
	data, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic("Invalid input format")
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Sprintf("Error parsing left number: %v", err))
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Sprintf("Error parsing left number: %v", err))
		}
		leftList = append(leftList, left)
		rightList = append(rightList, right)
		// fmt.Printf("Reading from file %v", leftList)
		// fmt.Printf("Reading from file %v", rightList)
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}

	return leftList, rightList
}
