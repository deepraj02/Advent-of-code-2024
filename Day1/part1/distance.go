package part1

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CalculateDistance(leftList, rightList []int) int {
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
