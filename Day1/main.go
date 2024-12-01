package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leftList, rightList := inputParser()
	totalDistance := calculateDistance(leftList, rightList)
	fmt.Printf("Total distance between lists: %d\n", totalDistance)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	if len(leftList) != len(rightList) {
		panic("Invalid input")
	}
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}
	return totalDistance
}

func inputParser() ([]int, []int) {
	data, err := os.OpenFile("inputs.txt", os.O_RDONLY, 0644)
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
		fmt.Printf("Reading from file %v", parts)
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
