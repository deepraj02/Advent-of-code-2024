package part2

func CalculateSimilarityScore(leftList, rightList []int) int {
	rightCount := make(map[int]int)
	for _, right := range rightList {
		rightCount[right]++
	}

	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}
	return similarityScore
}
